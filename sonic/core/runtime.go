package core

import (
	"context"
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/sonic/frame"
	"github.com/go-redis/redis/v8"
	jsoniter "github.com/json-iterator/go"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"
	"time"
)

var (
	json  = jsoniter.ConfigFastest
)



type RuntimeStats struct {
	Pod              string
	Deployment       string
	TasksRunning     int32  `json:"tasks_running"`
	MemoryAllocation string `json:"memory_allocation"`
	Goroutines       int    `json:"goroutines"`
}

func startRuntimeStats() {
	opts, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	rdb := redis.NewClient(opts)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		defer close(c)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		rdb.Del(context.Background(), fmt.Sprintf("runtime:channels:%s", os.Getenv("POD_NAME")))
		cancel()
	}()

	podName := os.Getenv("POD_NAME")
	if podName == "" {
		cancel()
		return
	}

	deploymentName := strings.Split(podName, "-")[0]

	rdb.Set(context.Background(), fmt.Sprintf("runtime:channels:%s", podName), fmt.Sprintf("%s:%s", deploymentName, podName), redis.KeepTTL)
	podType := os.Getenv("POD_TYPE")

	go func() {
		var m runtime.MemStats
		var stats = RuntimeStats{
			Pod:        podName,
			Deployment: deploymentName,
		}
		for range time.Tick(time.Second * 3) {
			if podType == "MODULE" {
				stats.TasksRunning = frame.Statistics.Running
			}
			runtime.ReadMemStats(&m)
			stats.MemoryAllocation = fmt.Sprintf("%d MBs", bToMb(m.Alloc))
			stats.Goroutines = runtime.NumGoroutine() - 2

			data, _ := json.Marshal(&stats)

			rdb.Publish(ctx, fmt.Sprintf("runtime:%s:%s", deploymentName, podName), data)
		}
	}()

	return
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
