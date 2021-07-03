package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Stripe holds the schema definition for the Stripe entity.
type Stripe struct {
	ent.Schema
}

// Fields of the Stripe.
func (Stripe) Fields() []ent.Field {
	return []ent.Field{
		field.String("CustomerID"),
		field.String("SubscriptionID").
			Optional(),
		field.Time("RenewalDate").
			Optional(),
	}
}

// Edges of the Stripe.
func (Stripe) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("License", License.Type).
			Ref("Stripe").
			Unique(),
	}
}
