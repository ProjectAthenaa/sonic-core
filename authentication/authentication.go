package authentication

import (
	"context"
	"fmt"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AuthenticationFunc(ctx context.Context) (context.Context, error) {
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

	return newCtx, nil
}

