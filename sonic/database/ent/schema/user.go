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
		field.Bool("Disabled").
			Default(false),
		field.Int("TasksRan").
			Default(0),
		field.Int("TotalDeclines").
			Default(0),
		field.Float("MoneySpent").
			Default(0.0),
		field.Int("TotalCheckouts").
			Default(0),
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
		edge.To("Checkouts", Checkout.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.NoAction,
			}),
		edge.To("App", App.Type).
			Unique().
			Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("Metadata", Metadata.Type).
			Unique(),
		edge.To("Sessions", Session.Type),
		edge.From("Release", Release.Type).
			Ref("Customers").
			Unique(),
	}
}
