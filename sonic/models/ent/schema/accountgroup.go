package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ProjectAthenaa/sonic-core/sonic"
)

// AccountGroup holds the schema definition for the AccountGroup entity.
type AccountGroup struct {
	ent.Schema
}

// Fields of the AccountGroup.
func (AccountGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),
		field.Enum("Site").Values(Sites...),
		field.Other("Accounts", sonic.Map{}).SchemaType(map[string]string{
			dialect.Postgres: "bytea",
		}),
	}
}

// Edges of the AccountGroup.
func (AccountGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("App", App.Type).
			Ref("AccountGroups").
			Unique(),
	}
}
