package frame

import (
	"context"
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/sonic/base"
	"github.com/ProjectAthenaa/sonic-core/sonic/core"
	"github.com/prometheus/common/log"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"
)

func init() {
	log.Info("Initializing runtime info streams")
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		defer close(c)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		cancel()
	}()

	podName := os.Getenv("POD_NAME")

	if podName == "" {
		cancel()
		return
	}

	deploymentName := strings.Split(podName, "-")[0]

	podType := os.Getenv("POD_TYPE")

	go func() {
		var m runtime.MemStats
		for range time.Tick(time.Second) {
			if podType == "MODULE" {
				core.Base.GetRedis("cache").Publish(ctx, fmt.Sprintf("runtime:%s:%s:tasks", deploymentName, podName), base.Statistics.Running)
			}

			runtime.ReadMemStats(&m)

			core.Base.GetRedis("cache").Publish(ctx, fmt.Sprintf("runtime:%s:%s:memory_allocation", deploymentName, podName), bToMb(m.Alloc))

			count := runtime.NumGoroutine() - 2
			core.Base.GetRedis("cache").Publish(ctx, fmt.Sprintf("runtime:%s:%s:goroutines", deploymentName, podName), count)
		}
	}()

}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
