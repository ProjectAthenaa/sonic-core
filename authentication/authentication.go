package authentication

import (
	"context"
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/sonic"
	"github.com/ProjectAthenaa/sonic-core/sonic/face"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
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

func GenGraphQLAuthenticationFunc(base face.ICoreContext, sessionCallback sessionCallback, noAuthResolvers ...interface{}) func() gin.HandlerFunc {
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

			ctx := context.WithValue(c.Request.Context(), "IP", c.Request.Header.Get("x-original-forwarded-for"))
			ctx = context.WithValue(ctx, "Location", c.Request.Header.Get("cf-ipcountry"))
			span := sentry.StartSpan(c.Request.Context(), "Authentication Middleware", sentry.TransactionName("Authentication"))
			defer span.Finish()

			if c.Request.URL.Path == "/query" {
				var body []byte
				c.Request.Body, body = sonic.NopCloserBody(c.Request.Body)

				c.JSON(200, gin.H{
					"errors": []map[string]interface{}{
						{
							"message": "error_shit",
							"path":    operationNameRe.FindStringSubmatch(string(body))[2],
						},
					},
					"data": nil,
				})

				//check if body contains no auth resolver
				if contains(string(body), noAuthResolverNames...) {
					goto setRequestContext
				}

				sessionID, err := c.Cookie("session_id")
				if err != nil && err != http.ErrNoCookie {
					c.JSON(403, unauthorizedError)
					return
				}

				if sessionID == "" {
					headerSession := c.GetHeader("Authorization")
					if strings.Contains(headerSession, "Bearer") {
						sessionID = strings.Split(headerSession, "Bearer ")[1]
					}

					user, err := extractTokensGin(base, c, sessionID)
					if err != nil {
						c.JSON(403, unauthorizedError)
						return
					}

					if user.IP != c.Request.Header.Get("x-original-forwarded-for") {
						c.JSON(428, ipDoesNotMatchSessionError)
						return
					}

					ctx = context.WithValue(ctx, "userID", user.UserID.String())
					ctx = context.WithValue(ctx, "discordID", user.DiscordID)
					goto setRequestContext
				}

				if sessionCallback != nil {
					ctx, err = sessionCallback(ctx, sessionID)
					if err != nil {
						c.JSON(403, err)
						return
					}
				}
			}

		setRequestContext:
			c.Request = c.Request.WithContext(ctx)
			c.Next()
		}
	}
}
