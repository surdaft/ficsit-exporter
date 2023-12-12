package exporter

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

type MetricVectorDetails struct {
	Name   string
	Help   string
	Labels []string
}

var RegisteredMetricVectors = []MetricVectorDetails{}
var RegisteredMetrics = []*prometheus.GaugeVec{}

func RegisterNewGaugeVec(reg *prometheus.Registry, opts prometheus.GaugeOpts, labelNames []string) *prometheus.GaugeVec {
	log.Printf("registering metric `%v`", opts)
	RegisteredMetricVectors = append(RegisteredMetricVectors, MetricVectorDetails{
		Name:   opts.Name,
		Help:   opts.Help,
		Labels: labelNames,
	})

	metric := prometheus.NewGaugeVec(opts, labelNames)
	err := reg.Register(metric)
	if err != nil {
		log.Printf("error registering metric `%s`", err.Error())
	}

	RegisteredMetrics = append(RegisteredMetrics, metric)
	return metric
}
