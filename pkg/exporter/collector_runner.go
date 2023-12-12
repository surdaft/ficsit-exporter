package exporter

import (
	"context"
	"sync"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type CollectorRunner struct {
	collectors []Collector
	ctx        context.Context
	cancel     context.CancelFunc
	registry   *prometheus.Registry
}

type Collector interface {
	Collect() error
}

func NewCollectorRunner(ctx context.Context, reg *prometheus.Registry, collectors ...Collector) *CollectorRunner {
	ctx, cancel := context.WithCancel(ctx)
	return &CollectorRunner{
		ctx:        ctx,
		cancel:     cancel,
		collectors: collectors,
		registry:   reg,
	}
}

func (c *CollectorRunner) Start() {
	errs := make(chan error)

	c.Collect(errs)
	for {
		select {
		case <-c.ctx.Done():
			return
		case <-Clock.After(5 * time.Second):
			c.Collect(errs)
		case <-errs:
			c.ctx.Done()
		}
	}
}

func (c *CollectorRunner) Stop() {
	c.cancel()
}

func (c *CollectorRunner) Collect(errs chan error) {
	wg := sync.WaitGroup{}

	for _, collector := range c.collectors {
		wg.Add(1)

		go func(co Collector) {
			err := co.Collect()

			go func(err error) {
				// cool, tell the success gauge something had failed
				if err != nil {
					errs <- err
				}
			}(err)

			wg.Done()
		}(collector)
	}

	wg.Wait()
}
