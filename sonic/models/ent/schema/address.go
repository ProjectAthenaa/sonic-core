package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.String("AddressLine"),
		field.String("AddressLine2").
			Optional(),
		field.String("Country"),
		field.String("State"),
		field.String("City"),
		field.String("ZIP"),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Shipping", Shipping.Type).
			Ref("ShippingAddress").
			Required(),
		edge.From("Shipping", Shipping.Type).
			Ref("BillingAddress"),
	}
}
