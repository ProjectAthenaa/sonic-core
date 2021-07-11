package authentication

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"os"
)

func AuthenticationFunc(ctx context.Context) (context.Context, error) {
	span := sentry.StartSpan(ctx, "Authentication", sentry.TransactionName("Authentication"))
	defer span.Finish()
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	userID, appID, err := extractTokens(ctx, token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, fmt.Sprint(err))
	}

	newCtx := context.WithValue(ctx, "userID", userID)
	newCtx = context.WithValue(newCtx, "appID", appID)
	//newCtx = context.WithValue(newCtx, "encryptionKey", encPassword)

	if os.Getenv("ENVIRONMENT") == "Production" {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			newCtx = context.WithValue(newCtx, "IP", md.Get("x-real-ip")[0])
		}
	}
	return newCtx, nil
}
