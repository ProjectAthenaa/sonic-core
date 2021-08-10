package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ProjectAthenaa/sonic-core/sonic"
	"github.com/google/uuid"
	"time"
)

var Sites = []string{
	"FinishLine",
	"JD_Sports",
	"YeezySupply",
	"Supreme",
	"Eastbay_US",
	"Champs_US",
	"Footaction_US",
	"Footlocker_US",
	"Bestbuy",
	"Pokemon_Center",
	"Panini_US",
	"Topss",
	"Nordstorm",
	"End",
	"Target",
	"Amazon",
	"Solebox",
	"Onygo",
	"Snipes",
	"Ssense",
	"Walmart",
	"Hibbet",
}

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("Name"),
		field.String("Image").
			Optional(),
		field.Enum("LookupType").
			Values("Keywords", "Link"),
		field.Strings("PositiveKeywords").Optional(),
		field.Strings("NegativeKeywords").Optional(),
		field.String("Link").Optional(),
		field.Int32("Quantity").Default(1),
		field.Strings("Sizes").Optional(),
		field.Strings("Colors").Optional(),
		field.Enum("Site").Values(Sites...),
		field.Other("Metadata", sonic.Map{}).SchemaType(map[string]string{dialect.Postgres: "bytea"}).Optional(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Task", Task.Type).
			Ref("Product").
			Required(),
		edge.From("Statistic", Statistic.Type).
			Ref("Product"),
		edge.From("Calendar", Calendar.Type).
			Ref("QuickTask").
			Unique(),
	}
}
