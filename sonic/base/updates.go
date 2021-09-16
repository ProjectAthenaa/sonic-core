package base

import (
	"context"
	"github.com/ProjectAthenaa/sonic-core/sonic/core"
	"sync"
	"time"
)

type update struct {
	id      string
	payload string
	channel string
}

type buffer map[string]update

var (
	updateLocker = &sync.Mutex{}
	updateBuffer = buffer{}
)

func (b buffer) queue(u update) {
	updateLocker.Lock()
	defer updateLocker.Unlock()
	b[u.id] = u
}

func init() {
	rdb := core.Base.GetRedis("cache")
	go func() {
		for range time.Tick(time.Millisecond * 200) {
			pipe := rdb.Pipeline()
			updateLocker.Lock()
			for _, v := range updateBuffer {
				pipe.Publish(context.Background(), v.channel, v.payload)
			}
			pipe.Exec(context.Background())
			updateLocker.Unlock()
		}
	}()
}
