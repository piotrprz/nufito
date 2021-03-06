package main

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
	"github.com/piotrprz/nufito/shared"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           shared.NufitoService
}

func (mw instrumentingMiddleware) GetTrainers() (output []shared.Trainer, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "getTrainers", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	output, err = mw.next.GetTrainers()
	return
}

func (mw instrumentingMiddleware) AddTrainer(trainer string) (err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "AddTrainer", "error", fmt.Sprint(err != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	err = mw.next.AddTrainer(trainer)
	return
}
