package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

var LicenseTypes = []string{
	"Unlocked",
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
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
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
