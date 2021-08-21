package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// ProfileGroup holds the schema definition for the ProfileGroup entity.
type ProfileGroup struct {
	ent.Schema
}

// Fields of the ProfileGroup.
func (ProfileGroup) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("Name"),
	}
}

// Edges of the ProfileGroup.
func (ProfileGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Profiles", Profile.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("App", App.Type).
			Ref("ProfileGroups").
			Required(),
		edge.From("Task", Task.Type).
			Ref("ProfileGroup").
			Unique(),
	}
}
