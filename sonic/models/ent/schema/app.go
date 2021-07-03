package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// App holds the schema definition for the App entity.
type App struct {
	ent.Schema
}

// Fields of the App.
func (App) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("first_login"),
	}
}

// Edges of the App.
func (App) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("User", User.Type).
			Ref("App").
			Unique().
			Required(),
		edge.To("Settings", Settings.Type),
		edge.To("ProxyLists", ProxyList.Type),
		edge.To("ProfileGroups", ProfileGroup.Type),
		edge.To("TaskGroups", TaskGroup.Type),
		edge.To("AccountGroups", AccountGroup.Type),
	}
}
