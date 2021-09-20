package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Checkout holds the schema definition for the Statistic entity.
type Checkout struct {
	ent.Schema
}

// Fields of the Checkout.
func (Checkout) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("Date").
			Default(time.Now),
		field.String("ProductName").
			Default("Unknown"),
		field.Float("ProductPrice").
			Default(0.0),
		field.String("ProductImage").
			Default("https://cdn.athenabot.com/default_product.svg"),
	}
}

// Edges of the Checkout.
func (Checkout) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("User", User.Type).
			Ref("Checkouts").
			Unique(),
	}
}
