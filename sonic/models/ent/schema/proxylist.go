package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// ProxyList holds the schema definition for the ProxyList entity.
type ProxyList struct {
	ent.Schema
}

// Fields of the ProxyList.
func (ProxyList) Fields() []ent.Field {
	return []ent.Field{
		field.String("Name"),
		field.Enum("Type").
			Values("Residential", "Datacenter", "ISP"),
	}
}

// Edges of the ProxyList.
func (ProxyList) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("App", App.Type).
			Ref("ProxyLists"),
		edge.To("Proxies", Proxy.Type),
		edge.From("Task", Task.Type).
			Ref("ProxyList"),
	}
}
