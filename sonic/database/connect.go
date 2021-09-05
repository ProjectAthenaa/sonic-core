package database

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/hook"
	_ "github.com/ProjectAthenaa/sonic-core/sonic/database/ent/runtime"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"os"
	"strings"
	"time"
)

func Connect(pgURL string) *ent.Client {
	client, err := ent.Open(dialect.Postgres, pgURL)
	if err != nil {
		panic(err)
	}

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
				return hook.AccountGroupFunc(func(ctx context.Context, m *ent.AccountGroupMutation) (ent.Value, error) {
					accounts, ok := m.Accounts()
					if !ok {
						return next.Mutate(ctx, m)
					}

					appID, _ := m.AppID()

					app := m.Client().App.GetX(ctx, appID)

					user, _ := app.User(ctx)

					site, _ := m.Site()

					setKey := fmt.Sprintf("accounts:%s:%s", strings.ToLower(string(site)))
					fmt.Println("creating account group: ", setKey)

					for username, password := range accounts {
						rdb.SAdd(ctx, setKey, user.ID.String(), fmt.Sprintf("%s:%s", username, password))
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

					appID, _ := m.AppID()
					app := m.Client().App.GetX(ctx, appID)
					user, _ := app.User(ctx)
					userID := user.ID.String()
					site, _ := m.Site()

					setKey := fmt.Sprintf("accounts:%s:%s", site, userID)

					var toDelete []string
					var toSet []string

					for newUsername, newPassword := range newAccounts {
						for oldUsername, oldPassword := range oldAccounts {
							if newUsername == oldUsername && newPassword != oldPassword {
								toDelete = append(toDelete, hash(fmt.Sprintf("%s:%s", newUsername, oldPassword)))
								toSet = append(toSet, fmt.Sprintf("%s:%s", newUsername, newPassword))
							} else if newUsername != oldUsername && newPassword == oldPassword {
								toDelete = append(toDelete, hash(fmt.Sprintf("%s:%s", oldUsername, oldPassword)))
								toSet = append(toSet, fmt.Sprintf("%s:%s", newUsername, newPassword))
							}
						}
					}

					for _, deletion := range toDelete {
						rdb.Set(ctx, deletion, "1", time.Hour*168)
					}

					for _, account := range toSet {
						rdb.SAdd(ctx, setKey, account)
					}

					return next.Mutate(ctx, m)
				})
			},
			ent.OpUpdateOne|ent.OpUpdate,
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
