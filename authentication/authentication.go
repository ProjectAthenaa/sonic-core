package authentication

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func AuthenticationFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "basic")
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

type CachedUser struct {
	UserID    uuid.UUID `json:"user_id"`
	LicenseID uuid.UUID `json:"key"`
	AppID     uuid.UUID `json:"app_id"`
	SessionID uuid.UUID `json:"session_id"`
	LoginTime int64 `json:"login_time"`
}
