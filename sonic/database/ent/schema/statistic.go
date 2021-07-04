package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Statistic holds the schema definition for the Statistic entity.
type Statistic struct {
	ent.Schema
}

// Fields of the Statistic.
func (Statistic) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Enum("Type").Values("Checkout", "Decline"),
	}
}

// Edges of the Statistic.
func (Statistic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("User", User.Type).
			Ref("Statistics"),
		edge.To("Product", Product.Type),
	}
}
