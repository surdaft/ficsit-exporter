package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	MachineItemsProducedPerMin   *prometheus.GaugeVec
	MachineItemsProducedEffiency *prometheus.GaugeVec
)

func RegisterFactoryMetrics(reg *prometheus.Registry) {
	MachineItemsProducedPerMin = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "machine_items_produced_per_min",
		Help: "How much of an item a building is producing",
	}, []string{
		"item_name",
		"machine_name",
		"x",
		"y",
		"z",
	})

	MachineItemsProducedEffiency = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "machine_items_produced_pc",
		Help: "The efficiency with which a building is producing an item",
	}, []string{
		"item_name",
		"machine_name",
		"x",
		"y",
		"z",
	})
}
