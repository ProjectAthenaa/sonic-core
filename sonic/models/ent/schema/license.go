package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

var LicenseTypes = []string{
	"Lifetime",
	"Renewal",
	"Beta",
	"Weekly",
	"FNF",
}

// License holds the schema definition for the License entity.
type License struct {
	ent.Schema
}

// Fields of the License.
func (License) Fields() []ent.Field {
	return []ent.Field{
		field.String("Key"),
		field.String("HardwareID").
			Optional(),
		field.String("MobileHardwareID").
			Optional(),
		field.Enum("Type").Values(LicenseTypes...),
	}
}

// Edges of the License.
func (License) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("User", User.Type).
			Ref("License").
			Unique().
			Required(),
		edge.To("Stripe", Stripe.Type),
	}
}
