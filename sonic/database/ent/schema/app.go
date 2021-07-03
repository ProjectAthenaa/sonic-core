package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

// Fields of the App.
func (App) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Bool("first_login").Default(false),
	}
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("User", User.Type).
			Ref("App").
			Unique().
			Required(),
		edge.To("Settings", Settings.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("ProxyLists", ProxyList.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("ProfileGroups", ProfileGroup.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("TaskGroups", TaskGroup.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.To("AccountGroups", AccountGroup.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
	}
}
