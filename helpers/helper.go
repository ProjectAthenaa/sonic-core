package helpers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/authentication"
	"github.com/ProjectAthenaa/sonic-core/sonic/core"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
	"github.com/google/uuid"
	"time"
)

func AuthenticateUser(user *ent.User) string {
	if user == nil {
		return ""
	}
	ctx := context.Background()
	c := authentication.CachedUser{
		UserID:    user.ID,
		LicenseID: user.QueryLicense().FirstX(ctx).ID,
		AppID:     user.QueryApp().FirstX(ctx).ID,
		SessionID: uuid.New(),
		LoginTime: time.Now().Add(time.Hour * 50000).Unix(),
	}

	rdb := core.Base.GetRedis("cache")

	val, _ := json.Marshal(c)

	key := "users:" + user.ID.String()
	rdb.Set(ctx, key, string(val), time.Minute)
	fmt.Println("set-login:", key)
	return user.ID.String()
}
