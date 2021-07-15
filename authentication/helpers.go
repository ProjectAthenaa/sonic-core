package authentication

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ProjectAthenaa/sonic-core/sonic"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/session"
	"github.com/ProjectAthenaa/sonic-core/sonic/face"
	"github.com/getsentry/sentry-go"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"google.golang.org/grpc/peer"
)

var (
	sessionExpiredError = errors.New("session_expired")
)

func extractTokens(base face.ICoreContext, ctx context.Context, sessionID string) (string, string, error) {
	val, err := base.GetRedis("default").Get(ctx, sessionID).Result()
	if err != redis.Nil && err != nil {
		p, _ := peer.FromContext(ctx)
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetContext("Context", map[string]interface{}{
				"IP":        p.Addr.String(),
				"SessionID": sessionID,
			})
		})
		return "", "", err
	}

	if val == "" || len(val) == 0 {
		_, err = base.GetPg("default").Session.Update().Where(session.ID(sonic.UUIDParser(sessionID))).SetExpired(true).Save(ctx)
		if err != nil {
			return "", "", err
		}
		return "", "", sessionExpiredError
	}

	var user CachedUser
	if err = json.Unmarshal([]byte(val), &user); err != nil {
		return "", "", err
	}

	return user.UserID.String(), user.AppID.String(), nil
}

type CachedUser struct {
	UserID    uuid.UUID `json:"user_id"`
	LicenseID uuid.UUID `json:"key"`
	AppID     uuid.UUID `json:"app_id"`
	SessionID uuid.UUID `json:"session_id"`
	DiscordID string    `json:"discord_id"`
	LoginTime int64     `json:"login_time"`
}
