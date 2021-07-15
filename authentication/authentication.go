package authentication

import (
	"context"
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/sonic"
	"github.com/ProjectAthenaa/sonic-core/sonic/face"
	"github.com/getsentry/sentry-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"os"
)

func GenAuthenticationFunc(base face.ICoreContext) func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		span := sentry.StartSpan(ctx, "Authentication", sentry.TransactionName("Authentication"))
		defer span.Finish()
		token, err := grpc_auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, err
		}

		userID, appID, err := extractTokens(base, ctx, token)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, fmt.Sprint(err))
		}

		newCtx := context.WithValue(ctx, "userID", userID)
		newCtx = context.WithValue(newCtx, "appID", appID)
		//newCtx = context.WithValue(newCtx, "encryptionKey", encPassword)

		if os.Getenv("ENVIRONMENT") == "Production" {
			newCtx = context.WithValue(newCtx, "IP", sonic.IPFromContext(ctx))
		}
		return newCtx, nil
	}
}
