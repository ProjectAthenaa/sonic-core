package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Bool("Disabled").Default(false),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("License", License.Type).
			Unique().
			Annotations(
				entsql.Annotation{
					OnDelete: entsql.Cascade,
				}),
		edge.To("Statistics", Statistic.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.NoAction,
			}),
		edge.To("App", App.Type).
			Unique().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
