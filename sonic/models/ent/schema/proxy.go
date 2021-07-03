package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Proxy holds the schema definition for the Proxy entity.
type Proxy struct {
	ent.Schema
}

// Fields of the Proxy.
func (Proxy) Fields() []ent.Field {
	return []ent.Field{
		field.String("Username").
			Optional(),
		field.String("Password").
			Optional(),
		field.String("IP"),
		field.String("Port"),
	}
}

// Edges of the Proxy.
func (Proxy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("ProxyList", ProxyList.Type).
			Ref("Proxies").
			Unique(),
	}
}
