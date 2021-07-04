package database

import (
	"context"
	"entgo.io/ent/dialect"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/migrate"
)

func Connect(pgURL string) *ent.Client {
	client, err := ent.Open(dialect.Postgres, pgURL)
	if err != nil {
		panic(err)
	}

	err = client.Schema.Create(context.Background(), migrate.WithDropIndex(true), migrate.WithDropColumn(true))
	if err != nil {
		panic(err)
	}
	return client
}
