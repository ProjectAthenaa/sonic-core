package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Shipping holds the schema definition for the Shipping entity.
type Shipping struct {
	ent.Schema
}

// Fields of the Shipping.
func (Shipping) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("FirstName"),
		field.String("LastName"),
		field.String("PhoneNumber"),
		field.Bool("BillingIsShipping"),
	}
}

// Edges of the Shipping.
func (Shipping) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Profile", Profile.Type).
			Ref("Shipping").
			Unique().
			Required(),
		edge.To("ShippingAddress", Address.Type),
		edge.To("BillingAddress", Address.Type),
	}
}
