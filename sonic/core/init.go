package core

import (
	"github.com/ProjectAthenaa/sonic-core/sonic/frame"
	"os"
)

var Base = &frame.CoreContext{}

func init() {
	_, _ = Base.NewPg("default", os.Getenv("PG_URL"))
	_, _ = Base.NewRedis("default", os.Getenv("REDIS_URL"))
}
