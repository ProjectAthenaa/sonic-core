package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Billing holds the schema definition for the Billing entity.
type Billing struct {
	ent.Schema
}

// Fields of the Billing.
func (Billing) Fields() []ent.Field {
	return []ent.Field{
		field.String("CardholderName"),
		field.String("CardNumber"),
		field.String("ExpiryMonth"),
		field.String("ExpiryYear"),
		field.String("CVV"),
	}
}

// Edges of the Billing.
func (Billing) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Profile", Profile.Type).
			Ref("Billing").
			Required(),
	}
}
