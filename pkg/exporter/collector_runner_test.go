package exporter_test

import (
	"context"

	"time"

	"github.com/AP-Hunt/ficsit-exporter/pkg/exporter"
	"github.com/benbjohnson/clock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/prometheus/client_golang/prometheus"
)

type TestCollector struct {
	counter int
}

func NewTestCollector(errs chan error) *TestCollector {
	return &TestCollector{
		counter: 0,
	}
}
func (t *TestCollector) Collect() error {
	t.counter = t.counter + 1
	return nil
}

var _ = Describe("CollectorRunner", func() {
	Describe("Basic Functionality", func() {
		It("runs on init and on each timeout", func() {
			errs := make(chan error)
			reg := prometheus.NewRegistry()
			ctx, cancel := context.WithCancel(context.Background())
			testTime := clock.NewMock()
			exporter.Clock = testTime

			c1 := NewTestCollector(errs)
			c2 := NewTestCollector(errs)
			runner := exporter.NewCollectorRunner(ctx, reg, c1, c2)
			go runner.Start()
			testTime.Add(5 * time.Second)
			testTime.Add(5 * time.Second)
			testTime.Add(5 * time.Second)
			cancel()
			Expect(c1.counter).To(Equal(3))
			Expect(c2.counter).To(Equal(3))
		})

		It("does not run after being canceled", func() {
			errs := make(chan error)
			reg := prometheus.NewRegistry()
			ctx, cancel := context.WithCancel(context.Background())
			testTime := clock.NewMock()
			exporter.Clock = testTime

			c1 := NewTestCollector(errs)
			runner := exporter.NewCollectorRunner(ctx, reg, c1)
			go runner.Start()
			testTime.Add(5 * time.Second)
			cancel()
			testTime.Add(5 * time.Second)
			testTime.Add(5 * time.Second)
			Expect(c1.counter).To(Equal(1))
		})
	})
})
