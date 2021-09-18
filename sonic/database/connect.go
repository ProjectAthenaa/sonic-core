package database

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/accountgroup"
	app2 "github.com/ProjectAthenaa/sonic-core/sonic/database/ent/app"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/hook"
	_ "github.com/ProjectAthenaa/sonic-core/sonic/database/ent/runtime"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	_ "github.com/lib/pq"
	"github.com/prometheus/common/log"
	"os"
	"regexp"
	"strings"
	"time"
)

var statusRe = regexp.MustCompile(`Status:(?:27|8|13|19)`)

func Connect(pgURL string) *ent.Client {
	client, err := ent.Open(dialect.Postgres, pgURL)
	if err != nil {
		panic(err)
	}

	redisSync := redsync.New(goredis.NewPool(rdb))

	client.Task.Use(
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.TaskFunc(func(ctx context.Context, mutation *ent.TaskMutation) (ent.Value, error) {
					if id, ok := mutation.ID(); ok {
						rdb.Publish(ctx, "scheduler:tasks-deleted", id.String())
					}
					return next.Mutate(ctx, mutation)
				})
			},
			ent.OpDeleteOne|ent.OpDelete,
		),
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.TaskFunc(func(ctx context.Context, mutation *ent.TaskMutation) (ent.Value, error) {
					id, _ := mutation.ID()
					oldST, _ := mutation.OldStartTime(ctx)
					product, err := mutation.Client().Product.Get(ctx, mutation.ProductIDs()[0])
					if err != nil {
						log.Error("error retrieving product: ", err)
						return next.Mutate(ctx, mutation)
					}

					newST, ok := mutation.StartTime()
					if !ok {
						rdb.ZRem(ctx, "zset:scheduled:items", fmt.Sprintf("%s:%s:%s", product.Site, id.String(), oldST))
						return next.Mutate(ctx, mutation)
					}

					if newST.Unix() != oldST.Unix() && newST.Sub(time.Now()) > 0 {
						rdb.ZAdd(ctx, "zset:scheduled:items", &redis.Z{
							Score:  float64(newST.Unix()),
							Member: fmt.Sprintf("%s:%s:%s", product.Site, id.String(), oldST),
						})
					}

					return next.Mutate(ctx, mutation)
				})
			},
			ent.OpUpdate|ent.OpUpdateOne,
		),
	)

	client.AccountGroup.Use(
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.AccountGroupFunc(func(ctx context.Context, m *ent.AccountGroupMutation) (ent.Value, error) {
					accounts, ok := m.Accounts()
					if !ok {
						return next.Mutate(ctx, m)
					}

					appID, _ := m.AppID()
					app := m.Client().App.GetX(ctx, appID)
					user, _ := app.User(ctx)
					site, _ := m.Site()

					setKey := fmt.Sprintf("accounts:%s:%s", strings.ToLower(string(site)), user.ID.String())
					fmt.Println("creating account group: ", setKey)

					for username, password := range accounts {
						rdb.SAdd(ctx, setKey, fmt.Sprintf("%s:%s", username, password)).Err()
					}

					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate,
		),
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.AccountGroupFunc(func(ctx context.Context, m *ent.AccountGroupMutation) (ent.Value, error) {
					oldAccounts, err := m.OldAccounts(ctx)
					if err != nil {
						return next.Mutate(ctx, m)
					}

					newAccounts, _ := m.Accounts()

					id, _ := m.ID()

					app := m.Client().App.Query().Where(app2.HasAccountGroupsWith(accountgroup.ID(id))).FirstX(ctx)
					user, _ := app.User(ctx)
					site, _ := m.Site()

					setKey := fmt.Sprintf("accounts:%s:%s", strings.ToLower(string(site)), user.ID.String())

					var toDelete []string
					var toSet []string

					for newUsername, newPassword := range newAccounts {
						for oldUsername, oldPassword := range oldAccounts {
							if newUsername == oldUsername && newPassword != oldPassword {
								toDelete = append(toDelete, fmt.Sprintf("%s:%s", newUsername, oldPassword))
								toSet = append(toSet, fmt.Sprintf("%s:%s", newUsername, newPassword))
							} else if newUsername != oldUsername && newPassword == oldPassword {
								toDelete = append(toDelete, fmt.Sprintf("%s:%s", oldUsername, oldPassword))
								toSet = append(toSet, fmt.Sprintf("%s:%s", newUsername, newPassword))
							}
						}
					}

					for _, deletion := range toDelete {
						rdb.Set(ctx, fmt.Sprintf("accounts:delete:%s", hash(deletion)), "1", time.Hour*168)
						rdb.SRem(ctx, setKey, deletion)
					}

					for _, account := range toSet {
						rdb.SAdd(ctx, setKey, account)
					}

					return next.Mutate(ctx, m)
				})
			},
			ent.OpUpdateOne|ent.OpUpdate,
		))

	client.Proxy.Use(
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.ProxyFunc(func(ctx context.Context, m *ent.ProxyMutation) (ent.Value, error) {
					plID, exists := m.ProxyListID()
					if !exists {
						log.Error("proxy list id not found")
						return next.Mutate(ctx, m)
					}

					key := fmt.Sprintf("tasks:proxies:%s", plID.String())

					locker := redisSync.NewMutex(key + ":locker")

					if err = locker.LockContext(ctx); err != nil {
						log.Error("error acquiring proxy mutex: ", err)
						return nil, err
					}

					defer func() {
						if ok, err := locker.UnlockContext(ctx); !ok || err != nil {
							log.Error("error unlocking proxy mutex: ", err)
						}
					}()

					//payload, err := json.Marshal(&module.Proxy{
					//	Username: m.,
					//	Password: nil,
					//	IP:       "",
					//	Port:     "",
					//})

					return next.Mutate(ctx, m)
				})

			},
			ent.OpCreate,
		),
	)

	//err = client.Schema.Create(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	return client
}

var rdb = func() *redis.Client {
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	return redis.NewClient(opts)
}()

func hash(text string) string {
	algorithm := sha1.New()
	algorithm.Write([]byte(text))
	return hex.EncodeToString(algorithm.Sum(nil))
}
