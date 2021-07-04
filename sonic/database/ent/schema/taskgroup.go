package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// TaskGroup holds the schema definition for the TaskGroup entity.
type TaskGroup struct {
	ent.Schema
}

// Fields of the TaskGroup.
func (TaskGroup) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("Name").Default("Default"),
	}
}

// Edges of the TaskGroup.
func (TaskGroup) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("App", App.Type).
			Ref("TaskGroups").
			Required(),
		edge.To("Tasks", Task.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
