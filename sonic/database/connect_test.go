package database

import (
	"context"
	"fmt"
	"testing"
)

//rebuild schema
func TestConnect(t *testing.T) {
	client := Connect("postgresql://postgres:postgres@localhost:5432/defaultdb?sslmode=disable")
	err := client.Schema.Create(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("success")
}
