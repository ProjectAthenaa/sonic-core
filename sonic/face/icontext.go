package face

import (
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent"
	"github.com/go-redis/redis/v8"
)

type ICoreContext interface {
	GetRedis(name string) redis.UniversalClient
	GetPg(name string) *ent.Client
}
