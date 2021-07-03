package database

import "github.com/ProjectAthenaa/sonic-core/sonic/database/ent"

func Connect(pgURL string) *ent.Client {
	client, err := ent.Open("postgres", pgURL)
	if err != nil {
		panic(err)
	}
	return client
}
