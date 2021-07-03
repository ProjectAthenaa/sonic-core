package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TaskGroup holds the schema definition for the TaskGroup entity.
type TaskGroup struct {
	ent.Schema
}

// Fields of the TaskGroup.
func (TaskGroup) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),
	}
}

// Edges of the TaskGroup.
func (TaskGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("App", App.Type).
			Ref("TaskGroups").
			Required(),
		edge.To("Tasks", Task.Type),
	}
}
