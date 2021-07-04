package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Profile holds the schema definition for the Profile entity.
type Profile struct {
	ent.Schema
}

// Fields of the Profile.
func (Profile) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("Name"),
		field.String("Email"),
	}
}

// Edges of the Profile.
func (Profile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ProfileGroup", ProfileGroup.Type).
			Ref("Profiles").
			Unique().
			Required(),
		edge.To("Shipping", Shipping.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("Billing", Billing.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
