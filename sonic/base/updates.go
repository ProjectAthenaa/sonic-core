package base

import (
	"context"
	"github.com/ProjectAthenaa/sonic-core/sonic/core"
	"github.com/go-redis/redis/v8"
	"time"
)


var (
	updateBuffer redis.Pipeliner
)

func init() {
	rdb := core.Base.GetRedis("cache")
	updateBuffer = rdb.Pipeline()
	go func() {
		for range time.Tick(time.Millisecond * 200) {
			updateBuffer.Exec(context.Background())
		}
	}()
}
