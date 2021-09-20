package core

import (
	"context"
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/logs"
	"github.com/ProjectAthenaa/sonic-core/protos/module"
	monitorController "github.com/ProjectAthenaa/sonic-core/protos/monitorController"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"google.golang.org/grpc"
	"os"
	"os/signal"
	"syscall"
)

var (
	redisSync *redsync.Redsync
	rdb       redis.UniversalClient
)

func ListenAndServe(module string, server module.ModuleServer) {
	
	defer func() {
		if a := recover(); a != nil {
			log.Warnf("[server] [Recovered] [%s]", fmt.Sprint(a))
			ListenAndServe(module, server)
		}
	}()

	initializeVariables()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()
	log.Infof("[server] [%s] [Module Initialized]", module)

	for {
		select {
		case <-c:
			log.Info("[server] [SIGTERM invoked]")
			cancel()
		case <-ctx.Done():
			log.Info("[server] [ctx deadline exceeded]")
			return
		case task := <-tasksListener(ctx, module):
			go processTask(ctx, task, server)
			log.Infof("[server] [New Task Received] [%s]", task)
		}
	}
}

func initializeVariables() {
	rdb = Base.GetRedis("cache")
	pool := goredis.NewPool(rdb)
	redisSync = redsync.New(pool)

	conn, err := grpc.Dial("monitor-controller.general.svc.cluster.local:3000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("error dialing monitor-controller.general.svc.cluster.local:3000: ", err)
	}

	monitorClient = monitorController.NewMonitorClient(conn)
}
