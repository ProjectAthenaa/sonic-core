package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Calendar holds the schema definition for the Calendar entity.
type Calendar struct {
	ent.Schema
}

// Fields of the Calendar.
func (Calendar) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("ReleaseDate"),
		field.String("ProductImage").
			Default("placeholder_image_here"),
		field.String("ProductName"),
		field.Bool("HypedRelease"),
		field.Int("UsersRunning"),
	}
}

// Edges of the Calendar.
func (Calendar) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("QuickTask", Product.Type).
			Unique(),
	}
}
