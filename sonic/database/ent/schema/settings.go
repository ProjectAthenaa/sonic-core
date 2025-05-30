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

// Settings holds the schema definition for the Settings entity.
type Settings struct {
	ent.Schema
}

// Fields of the Settings.
func (Settings) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.String("SuccessWebhook").Default(""),
		field.String("DeclineWebhook").Default(""),
		field.Int32("CheckoutDelay").Default(0),
		field.Int32("ATCDelay").Default(0),
		field.Enum("CaptchaSolver").
			Values("AYCD_Autosolve", "2Captcha", "CapMonster", "Harvester", "Disabled").
			Default("Disabled"),
		field.Other("CaptchaDetails", sonic.Map{}).SchemaType(map[string]string{dialect.Postgres: "bytea"}),
	}
}

// Edges of the Settings.
func (Settings) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("App", App.Type).
			Ref("Settings").
			Required().
			Unique(),
	}
}
