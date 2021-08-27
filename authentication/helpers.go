package authentication

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ProjectAthenaa/sonic-core/sonic"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/session"
	"github.com/ProjectAthenaa/sonic-core/sonic/face"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"google.golang.org/grpc/peer"
	"index/suffixarray"
	"regexp"
	"strings"
	"time"
)

type sessionCallback = func(ctx context.Context, sessionID string) (context.Context, error)

var (
	sessionExpiredError = errors.New("session_expired")
	operationNameRe     = regexp.MustCompile(`(mutation|query|subscription)\s+{\s+(\w+)?`)
)

func extractTokens(base face.ICoreContext, ctx context.Context, sessionID string) (string, string, string, error) {
	val, err := base.GetRedis("cache").Get(ctx, "users:"+sessionID).Result()
	if err != redis.Nil && err != nil {
		p, _ := peer.FromContext(ctx)
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetContext("Context", map[string]interface{}{
				"IP":        p.Addr.String(),
				"SessionID": sessionID,
			})
		})
		return "", "", "", err
	}

	if val == "" || len(val) == 0 {
		_, err = base.GetPg("default").Session.Update().Where(session.ID(sonic.UUIDParser(sessionID))).SetExpired(true).Save(ctx)
		if err != nil {
			return "", "", "", err
		}
		return "", "", "", sessionExpiredError
	}

	var user CachedUser
	if err = json.Unmarshal([]byte(val), &user); err != nil {
		return "", "", "", err
	}

	if time.Since(user.LastRefresh) >= time.Minute*30 || time.Since(user.LoginTime) >= time.Hour*24 {
		base.GetRedis("cache").Del(ctx, "users:"+sessionID)
		return "", "", "", errors.New("session_expired")
	}

	return user.UserID.String(), user.AppID.String(), user.IP, nil
}

func extractTokensGin(base face.ICoreContext, ctx *gin.Context, sessionID string) (*CachedUser, error) {
	sessionID = "users:" + sessionID
	val, err := base.GetRedis("cache").Get(ctx, sessionID).Result()
	if err != redis.Nil && err != nil {
		sentry.WithScope(func(scope *sentry.Scope) {
			scope.SetContext("Context", map[string]interface{}{
				"IP":        ctx.Request.Header.Get("x-original-forwarded-for"),
				"SessionID": sessionID,
			})
		})
		return nil, err
	}

	if val == "" || len(val) == 0 {
		_, err = base.GetPg("pg").Session.Update().Where(session.ID(sonic.UUIDParser(sessionID))).SetExpired(true).Save(ctx)
		if err != nil {
			return nil, err
		}
		return nil, sessionExpiredError
	}

	var user CachedUser
	if err = json.Unmarshal([]byte(val), &user); err != nil {
		return nil, err
	}

	return &user, nil
}

type CachedUser struct {
	UserID      uuid.UUID `json:"user_id"`
	LicenseID   uuid.UUID `json:"key"`
	AppID       uuid.UUID `json:"app_id"`
	SessionID   uuid.UUID `json:"session_id"`
	DiscordID   string    `json:"discord_id"`
	LoginTime   time.Time `json:"login_time"`
	LastRefresh time.Time `json:"last_refresh"`
	HardwareID  string    `json:"hardware_id"`
	IP          string    `json:"ip"`
	Mobile      bool      `json:"mobile"`
}

type GraphQLError struct {
	Errors []struct {
		Message string   `json:"message"`
		Path    []string `json:"path"`
	} `json:"errors"`
	Data interface{} `json:"data"`
}

func contains(body string, subStrs ...string) bool {
	if strings.Contains(body, "IntrospectionQuery") {
		return true
	}

	var str string
	if v := operationNameRe.FindStringSubmatch(body); len(v) >= 2 {
		str = v[2]
	} else {
		return false
	}
	if len(subStrs) == 0 {
		return true
	}
	r := regexp.MustCompile("(?i)" + strings.Join(subStrs, "|"))
	index := suffixarray.New([]byte(str))
	res := index.FindAllIndex(r, -1)
	exists := make(map[string]int)
	for _, v := range subStrs {
		exists[v] = 1
	}
	for _, pair := range res {
		s := str[pair[0]:pair[1]]
		exists[s] = exists[s] + 1
	}
	for _, v := range exists {
		if v == 1 {
			return false
		}
	}
	return true
}

func ExtractFromCtx(ctx context.Context) (*uuid.UUID, error) {
	if e := ctx.Value("error"); e != nil {
		return nil, e.(error)
	}

	if userID := ctx.Value("userID"); userID != nil {
		id := userID.(uuid.UUID)
		return &id, nil
	}

	return nil, errors.New("user_not_found")
}
