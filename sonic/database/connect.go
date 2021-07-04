package database

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
)

func Connect(pgURL string) *ent.Client {
	client, err := ent.Open(dialect.Postgres, pgURL)
	if err != nil {
		panic(err)
	}

	err = client.Schema.Create(context.Background())
	if err != nil {
		panic(err)
	}
	return client
}
