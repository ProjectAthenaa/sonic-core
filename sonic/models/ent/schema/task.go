package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Time("StartTime").
			Optional(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Product", Product.Type),
		edge.To("ProxyList", ProxyList.Type),
		edge.To("ProfileGroup", ProfileGroup.Type),
		edge.From("TaskGroup", TaskGroup.Type).
			Ref("Tasks"),
	}
}
