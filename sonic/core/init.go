package core

import (
	"github.com/ProjectAthenaa/sonic-core/sonic/frame"
	"os"
)

var Base = &frame.CoreContext{}

func init() {
	_ = os.Setenv("PG_URL", "postgresql://postgres:postgres@127.0.0.1:5432/defaultdb?sslmode=disable")
	_ = os.Setenv("REDIS_URL", "rediss://default:n6luoc78ac44pgs0@test-redis-do-user-9223163-0.b.db.ondigitalocean.com:25061")
	_, _ = Base.NewPg("default", os.Getenv("PG_URL"))
	_, _ = Base.NewRedis("default", os.Getenv("REDIS_URL"))
}
