package base

import (
	"context"
	"github.com/ProjectAthenaa/sonic-core/sonic/core"
	"sync"
	"time"
)

type update struct {
	channel string
	payload string
}

var (
	updateMutex  = &sync.Mutex{}
	updateBuffer = make([]update, 0)
)

func queueUpdate(u update) {
	updateMutex.Lock()
	defer updateMutex.Unlock()
	updateBuffer = append(updateBuffer, u)
}

func init() {
	rdb := core.Base.GetRedis("cache")
	ctx := context.Background()
	go func() {
		for range time.Tick(time.Millisecond * 100) {
			pipe := rdb.Pipeline()

			for _, updateStatus := range updateBuffer {
				pipe.Publish(ctx, updateStatus.channel, updateStatus.payload)
			}
			pipe.Exec(ctx)
		}
	}()
}
