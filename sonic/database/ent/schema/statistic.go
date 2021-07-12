package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	gql "github.com/ProjectAthenaa/sonic-core/stats-graph/model"
	"github.com/google/uuid"
	"time"
)

// Statistic holds the schema definition for the Statistic entity.
type Statistic struct {
	ent.Schema
}

type Axis string

const (
	X Axis = "X"
	Y Axis = "Y"
)

// Fields of the Statistic.
func (Statistic) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Enum("Type").Values(
			string(gql.StatTypeCheckouts),
			string(gql.StatTypeDeclines),
			string(gql.StatTypeErrors),
			string(gql.StatTypeFailed),
			string(gql.StatTypeCookieGens),
			string(gql.StatTypeRecaptchaUsage),
			string(gql.StatTypeTasksRunning),
			string(gql.StatTypeMoneySpent),
		),
		field.Int("PotentialProfit").
			Optional().
			Nillable(),
		field.JSON("Axis", map[Axis]string{}),
		field.Int("Value").Optional(),
		field.Float("Spent").
			Default(0).
			Optional(),
	}
}

// Edges of the Statistic.
func (Statistic) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("User", User.Type).
			Ref("Statistics"),
		edge.To("Product", Product.Type),
	}
}
