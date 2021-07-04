package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// ProxyList holds the schema definition for the ProxyList entity.
type ProxyList struct {
	ent.Schema
}

// Fields of the ProxyList.
func (ProxyList) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("Name").Default("Default"),
		field.Enum("Type").
			Values("Residential", "Datacenter", "ISP"),
	}
}

// Edges of the ProxyList.
func (ProxyList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("App", App.Type).
			Ref("ProxyLists"),
		edge.To("Proxies", Proxy.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}),
		edge.From("Task", Task.Type).Annotations(entsql.Annotation{OnDelete: entsql.Cascade}).
			Ref("ProxyList"),
	}
}
