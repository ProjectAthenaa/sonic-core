package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
	"github.com/ProjectAthenaa/sonic-core/sonic"
	"github.com/google/uuid"
)

// Device holds the schema definition for the Device entity.
type Device struct {
	ent.Schema
}

// Fields of the Device.
func (Device) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("gpuVendor"),
		field.Strings("plugins"),
		field.Other("adevice", sonic.Map{}).SchemaType(map[string]string{dialect.Postgres: "bytea"}).Optional().Nillable(),
	}
}

// Edges of the Device.
func (Device) Edges() []ent.Edge {
	return nil
}


