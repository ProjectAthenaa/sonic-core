package database

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/hook"
	_ "github.com/ProjectAthenaa/sonic-core/sonic/database/ent/runtime"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"os"
)

func Connect(pgURL string) *ent.Client {
	client, err := ent.Open(dialect.Postgres, pgURL)
	if err != nil {
		panic(err)
	}

	client.Use(
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
