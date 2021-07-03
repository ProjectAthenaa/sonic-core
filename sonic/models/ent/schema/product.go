package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ProjectAthenaa/sonic-core/sonic"
	"github.com/lib/pq"
)

var Sites = []string{
	"FinishLine",
	"JD Sports",
	"YeezySupply",
	"Supreme",
	"Eastbay US",
	"Champs US",
	"Footaction US",
	"Footlocker US",
	"Bestbuy",
	"Pokemon Center",
	"Panini US",
	"Topss",
	"Nordstorm",
	"End.",
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
		field.String("Name"),
		field.String("Image").
			Optional(),
		field.Enum("LookupType").
			Values("Keywords", "Link"),
		field.Other("PositiveKeywords", pq.StringArray{}).Optional(),
		field.Other("NegativeKeywords", pq.StringArray{}).Optional(),
		field.String("Link").Optional(),
		field.Int("Quantity"),
		field.Other("Sizes", pq.StringArray{}),
		field.Other("Colors", pq.StringArray{}),
		field.Enum("Site").Values(Sites...),
		field.Other("Metadata", sonic.Map{}),
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
	}
}
