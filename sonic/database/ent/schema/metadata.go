package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Metadata holds the schema definition for the Metadata entity.
type Metadata struct {
	ent.Schema
}

// Fields of the Metadata.
func (Metadata) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Bool("FirstLogin").
			Default(true),
		field.Enum("Theme").
			Values("Variation1", "Variation2", "Variation3", "Variation4").
			Default("Variation1"),
	}
}

// Edges of the Metadata.
func (Metadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("Metadata").
			Unique().
			Required(),
	}
}
