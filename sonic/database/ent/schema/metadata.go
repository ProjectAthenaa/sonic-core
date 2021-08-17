package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"time"
)

// Metadata holds the schema definition for the Metadata entity.
type Metadata struct {
	ent.Schema
}

// Fields of the Metadata.
func (Metadata) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Bool("FirstLogin").
			Default(true),
		field.Bool("FirstLoginMobile").
			Default(true),
		field.Enum("Theme").
			Values("Variation1", "Variation2", "Variation3", "Variation4").
			Default("Variation1"),
		field.String("DiscordID").Default(""),
		field.String("DiscordAccessToken").Default(""),
		field.String("DiscordRefreshToken").Default(""),
		field.String("DiscordUsername").Default(""),
		field.String("DiscordAvatar").Default("https://cdn.athenabot.com/default_avatar.png"),
		field.String("DiscordDiscriminator").Default(""),
		field.Time("DiscordExpiryTime").Default(time.Now),
	}
}

// Edges of the Metadata.
func (Metadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("Metadata").
			Unique().
			Required(),
	}
}
