package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProfileGroup holds the schema definition for the ProfileGroup entity.
type ProfileGroup struct {
	ent.Schema
}

// Fields of the ProfileGroup.
func (ProfileGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),
	}
}

// Edges of the ProfileGroup.
func (ProfileGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Profiles", Profile.Type),
		edge.From("App", App.Type).
			Ref("ProfileGroups").
			Required(),
		edge.From("Task", Task.Type).
			Ref("ProfileGroup"),
	}
}
