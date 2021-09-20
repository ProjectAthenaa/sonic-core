package core

import (
	"context"
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/logs"
	"github.com/ProjectAthenaa/sonic-core/protos/module"
	monitorController "github.com/ProjectAthenaa/sonic-core/protos/monitorController"
	"github.com/evalphobia/logrus_sentry"
	"github.com/getsentry/raven-go"
	"github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/sirupsen/logrus"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
	"google.golang.org/grpc"
	"log/syslog"
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

	client, err := raven.New(os.Getenv("SENTRY_DSN"))
	if err != nil {
		log.Errorf("[server] [error initializing sentry] [%s]", fmt.Sprint(err))
	}

	hook, err := logrus_sentry.NewWithClientSentryHook(client, []logrus.Level{
		logrus.PanicLevel,
		logrus.FatalLevel,
		logrus.ErrorLevel,
	})

	if err == nil {
		log.AddHook(hook)
	}
	sysLogHook, err := logrus_syslog.NewSyslogHook("udp", "logs4.papertrailapp.com:44377", syslog.LOG_ERR|syslog.LOG_WARNING|syslog.LOG_NOTICE|syslog.LOG_INFO|syslog.LOG_DEBUG|, "")
	log.AddHook(sysLogHook)

}
