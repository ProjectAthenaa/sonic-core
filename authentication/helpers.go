package authentication

import (
	"context"
	"encoding/base64"
	"errors"
	"github.com/ProjectAthenaa/sonic-core/sonic/database"
	"github.com/google/uuid"
	"os"
	"strings"
)

var client = database.Connect(os.Getenv("PG_URL"))

func extractTokens(ctx context.Context, basicToken string) (string, string, error) {
	data, err := base64.StdEncoding.DecodeString(basicToken)
	if err != nil {
		return "", "", err
	}

	tokens := strings.Split(string(data), ":")

	if len(tokens) != 2 {
		return "", "", errors.New("BAD_TOKEN")
	}

	id, _ := uuid.Parse(tokens[0])

	user := client.User.GetX(ctx, id)

	if user.QueryLicense().FirstX(ctx).ID.String() == "" {
		return "", "", errors.New("NO_LICENSE")
	}

	if user.Disabled {
		return "", "", errors.New("USER_DISABLED")
	}

	if user.QueryApp().FirstX(ctx).ID.String() != tokens[1] {
		return "", "", errors.New("BAD_TOKEN")
	}

	return tokens[0], tokens[1], nil
}
