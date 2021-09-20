package core

import (
	"context"
	monitorController "github.com/ProjectAthenaa/sonic-core/protos/monitorController"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/product"
	"github.com/ProjectAthenaa/sonic-core/logs"
)

var monitorClient monitorController.MonitorClient

func (j *scratchTask) processMonitor(ctx context.Context) {
	log.Infof("[server] [starting monitor] [%s]", j.ID.String())
	newMonitorTask := &monitorController.Task{
		Site:         string(j.Edges.Product[0].Site),
		Metadata:     j.Edges.Product[0].Metadata,
		RedisChannel: j.getMonitorID(),
	}

	switch j.Edges.Product[0].LookupType {
	case product.LookupTypeKeywords:
		newMonitorTask.Lookup = &monitorController.Task_Keywords{Keywords: &monitorController.Keywords{
			Positive: j.Edges.Product[0].PositiveKeywords,
			Negative: j.Edges.Product[0].NegativeKeywords,
		}}
	case product.LookupTypeLink:
		newMonitorTask.Lookup = &monitorController.Task_Link{Link: j.Edges.Product[0].Link}
	case product.LookupTypeOther:
		break
	}

	resp, err := monitorClient.NewTask(ctx, newMonitorTask)
	if err != nil {
		log.Errorf("[server] [error starting monitor] [%s] [%s]", j.ID.String(), err)
		return
	}

	if resp.Stopped == true && resp.Error == nil {
		log.Warnf("[server] [monitor did not start] [%s]", j.ID.String())
		return
	} else if resp.Error != nil {
		log.Infof("[server] [%s] [%s]", *resp.Error, j.ID.String())
	}
}
