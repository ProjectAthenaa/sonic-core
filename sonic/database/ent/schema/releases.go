package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/thoas/go-funk"
	"time"
)

// Release holds the schema definition for the Release entity.
type Release struct {
	ent.Schema
}

// Fields of the Release.
func (Release) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("ReleaseDate").
			Default(time.Now),
		field.Int32("StockLevel").
			Default(0),
		field.String("Code").
			Default(funk.RandomString(10)),
		field.Enum("Type").
			Values(LicenseTypes...).
			Default("Renewal"),
		field.Int64("OneTimeFeeAmount").
			Default(60000),
		field.Int64("SubscriptionFee").
			Default(10000).
			Optional().
			Nillable(),
		field.String("PriceID").
			Optional().
			Nillable(),
	}
}

// Edges of the Metadata.
func (Release) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Customers", User.Type),
	}
}
