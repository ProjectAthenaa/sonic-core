package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Statistic holds the schema definition for the Statistic entity.
type Statistic struct {
	ent.Schema
}

// Fields of the Statistic.
func (Statistic) Fields() []ent.Field {
	return []ent.Field{
		field.Enum("Type").Values("Checkout", "Decline"),
	}
}

// Edges of the Statistic.
func (Statistic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("statistics"),
		edge.To("Product", Product.Type),
	}
}
