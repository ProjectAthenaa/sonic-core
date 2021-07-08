package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Session holds the schema definition for the Session entity.
type Session struct {
	ent.Schema
}

// Fields of the Session.
func (Session) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.String("DeviceName").
			Default("Unknown"),
		field.String("OS").
			Default("Unknown"),
		field.Enum("DeviceType").
			Values("Unknown", "Phone", "Tablet", "PC", "Laptop").
			Default("Unknown"),
		field.String("IP").Default("Unknown"),
	}
}

// Edges of the Metadata.
func (Session) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("Sessions").
			Unique(),
	}
}
