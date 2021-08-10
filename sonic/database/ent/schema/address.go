package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Address holds the schema definition for the Address entity.
type Address struct {
	ent.Schema
}

// Fields of the Address.
func (Address) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("AddressLine"),
		field.String("AddressLine2").
			Optional(),
		field.String("Country"),
		field.String("State"),
		field.String("City"),
		field.String("ZIP"),
		field.String("StateCode").
			Optional(),
	}
}

// Edges of the Address.
func (Address) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ShippingAddress", Shipping.Type).
			Ref("ShippingAddress").
			Required().
			Unique(),
		edge.From("BillingAddress", Shipping.Type).
			Ref("BillingAddress"),
	}
}
