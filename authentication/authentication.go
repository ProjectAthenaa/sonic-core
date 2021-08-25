package authentication

import (
	"context"
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/sonic"
	"github.com/ProjectAthenaa/sonic-core/sonic/face"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	"os"
	"reflect"
	"runtime"
	"strings"
)

func GenAuthenticationFunc(base face.ICoreContext) func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		span := sentry.StartSpan(ctx, "Authentication", sentry.TransactionName("Authentication"))
		defer span.Finish()
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "no token")
		}

		userID, appID, ip, err := extractTokens(base, ctx, token)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, fmt.Sprint(err))
		}

		newCtx := context.WithValue(ctx, "userID", userID)
		newCtx = context.WithValue(newCtx, "appID", appID)

		if os.Getenv("ENVIRONMENT") == "Production" {
			if ip != sonic.IPFromContext(ctx) {
				return nil, status.Error(codes.Unauthenticated, "ip_changed")
			}

			newCtx = context.WithValue(newCtx, "IP", sonic.IPFromContext(ctx))
		}
		return newCtx, nil
	}
}

func GenGraphQLAuthenticationFunc(base face.ICoreContext, graphEndpoint string, sessionCallback sessionCallback, noAuthResolvers ...interface{}) func() gin.HandlerFunc {
	var noAuthResolverNames []string

	for _, fun := range noAuthResolvers {
		if kind := reflect.TypeOf(fun).Kind(); kind == reflect.Func {
			noAuthResolverNames = append(noAuthResolverNames, runtime.FuncForPC(reflect.ValueOf(fun).Pointer()).Name())
		} else if kind == reflect.String {
			noAuthResolverNames = append(noAuthResolverNames, fun.(string))
		}
	}

	return func() gin.HandlerFunc {
		return func(c *gin.Context) {
			ip := c.Request.Header.Get("cf-connecting-ip")
			ctx := context.WithValue(c.Request.Context(), "IP", ip)
			ctx = context.WithValue(ctx, "Location", c.Request.Header.Get("cf-ipcountry"))
			span := sentry.StartSpan(c.Request.Context(), "Authentication Middleware", sentry.TransactionName("Authentication"))
			defer span.Finish()
			if strings.Contains(c.Request.URL.Path, graphEndpoint) {
				log.Info("Entered if statement")
				var body []byte
				c.Request.Body, body = sonic.NopCloserBody(c.Request.Body)
				//check if body contains no auth resolver
				if contains(string(body), noAuthResolverNames...) {
					goto setRequestContext
				}
				log.Info("Passed auth resolver exclusion check")

				sessionID, err := c.Cookie("session_id")
				if err != nil && err != http.ErrNoCookie {
					ctx = context.WithValue(ctx, "error", unauthorizedError)
					goto setRequestContext
				}
				log.Info("Retrieved session is from cookies")

				if sessionID == "" {
					headerSession := c.GetHeader("Authorization")
					if strings.Contains(headerSession, "Bearer") {
						sessionID = strings.Split(headerSession, "Bearer ")[1]
					}
				}
				log.Info("Retrieved session from headers")

				user, err := extractTokensGin(base, c, sessionID)
				if err != nil {
					ctx = context.WithValue(ctx, "error", unauthorizedError)
					goto setRequestContext
				}
				log.Info("Extracted user from session")

				log.Info(user.IP + " " + ip)
				log.Info(user.IP == ip)

				if user.IP != ip {
					ctx = context.WithValue(ctx, "error", ipDoesNotMatchSessionError)
					goto setRequestContext
				}
				log.Info("Passed ip check")

				ctx = context.WithValue(ctx, "userID", user.UserID)
				ctx = context.WithValue(ctx, "discordID", user.DiscordID)
				log.Info("Added everything to context")
				goto setRequestContext
			}

			//if sessionCallback != nil {
			//	ctx, err = sessionCallback(ctx, sessionID)
			//	if err != nil {
			//		ctx = context.WithValue(ctx, "error", err)
			//		goto setRequestContext
			//	}
			//}

		setRequestContext:
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		}
	}
}
