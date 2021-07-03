package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Settings holds the schema definition for the Settings entity.
type Settings struct {
	ent.Schema
}

// Fields of the Settings.
func (Settings) Fields() []ent.Field {
	return []ent.Field{
		field.String("SuccessWebhook"),
		field.String("DeclineWebhook"),
		field.Int("CheckoutDelay"),
		field.Int("ATCDelay"),
	}
}

// Edges of the Settings.
func (Settings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("App", App.Type).
			Ref("Settings").
			Required().
			Unique(),
	}
}
