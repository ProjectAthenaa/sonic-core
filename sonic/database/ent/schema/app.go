package schema

import (
	"entgo.io/ent"
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
		edge.To("Settings", Settings.Type),
		edge.To("ProxyLists", ProxyList.Type),
		edge.To("ProfileGroups", ProfileGroup.Type),
		edge.To("TaskGroups", TaskGroup.Type),
		edge.To("AccountGroups", AccountGroup.Type),
	}
}
