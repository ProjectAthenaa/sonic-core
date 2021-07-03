package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),
		field.String("Email"),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ProfileGroup", ProfileGroup.Type).
			Ref("Profile").
			Unique().
			Required(),
		edge.To("Shipping", Shipping.Type),
		edge.To("Billing", Billing.Type),
	}
}
