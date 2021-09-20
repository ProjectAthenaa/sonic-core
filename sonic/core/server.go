package core

import (
	"context"
	"fmt"
	"github.com/ProjectAthenaa/sonic-core/logs"
	"github.com/ProjectAthenaa/sonic-core/protos/module"
	"time"
)

func tasksListener(ctx context.Context, key string) <-chan string {
	tasks := make(chan string)
	key = fmt.Sprintf("queue:%s", key)
	go func() {
		defer func() {
			if a := recover(); a != nil {
				log.Warnf("[tasks listener] [recovered] [%s]", fmt.Sprint(a))
			}
		}()
		for {
			newTask := rdb.BLPop(ctx, time.Second, key).Val()

			if len(newTask) > 1 {
				tasks <- newTask[1]
			}
		}
	}()

	return tasks
}

func processTask(ctx context.Context, taskID string, server module.ModuleServer) {
	defer log.Infof("[server] [Task Processed] [%s]", taskID)
	data, err := getPayload(ctx, taskID)
	if err != nil {
		log.Errorf("[server] [error retrieving payload] [%s]", taskID)
		return
	}

	if data == nil {
		log.Errorf("[server] [payload is empty] [%s]", taskID)
		return
	}

	if _, err = server.Task(ctx, data); err != nil {
		log.Errorf("[server] [error starting task] [%s] [%s] ", taskID, err)
	}
}
