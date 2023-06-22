package jobs

import (
	"context"
	"route256/libs/jobs"
	"route256/loms/internal/config"
	"route256/loms/internal/service"
	"time"
)

type Jobs struct {
	svc  *service.Service
	jobs []jobs.Job
}

func NewJobs(svc *service.Service) *Jobs {
	return &Jobs{
		svc: svc,
	}
}

func (j *Jobs) Start(ctx context.Context) {
	cancelOrderJob := jobs.New(j.svc.CancelExpiredOrders)
	interval := time.Millisecond * time.Duration(config.AppConfig.Jobs.CancelOrder.Interval)
	cancelOrderJob.Start(ctx, interval)

	j.jobs = append(j.jobs, cancelOrderJob)
}

func (j *Jobs) Close() {
	for _, job := range j.jobs {
		job.Stop()
	}
}
