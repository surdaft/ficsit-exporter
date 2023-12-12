package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	ItemProductionCapacityPerMinute  *prometheus.GaugeVec
	ItemProductionCapacityPercent    *prometheus.GaugeVec
	ItemConsumptionCapacityPerMinute *prometheus.GaugeVec
	ItemConsumptionCapacityPercent   *prometheus.GaugeVec
	ItemsProducedPerMin              *prometheus.GaugeVec
	ItemsConsumedPerMin              *prometheus.GaugeVec
	PowerConsumed                    *prometheus.GaugeVec
	PowerCapacity                    *prometheus.GaugeVec
	PowerMaxConsumed                 *prometheus.GaugeVec
	BatteryDifferential              *prometheus.GaugeVec
	BatteryPercent                   *prometheus.GaugeVec
	BatteryCapacity                  *prometheus.GaugeVec
	BatterySecondsEmpty              *prometheus.GaugeVec
	BatterySecondsFull               *prometheus.GaugeVec
	FuseTriggered                    *prometheus.GaugeVec
)

func RegisterMetrics(reg *prometheus.Registry) {
	ItemProductionCapacityPerMinute = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "item_production_capacity_per_min",
		Help: "The factory's capacity for the production of an item, per minute",
	}, []string{
		"item_name",
	})

	ItemProductionCapacityPercent = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "item_production_capacity_pc",
		Help: "The percentage of an item's production capacity being used",
	}, []string{
		"item_name",
	})

	ItemConsumptionCapacityPerMinute = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "item_consumption_capacity_per_min",
		Help: "The factory's capacity for the consumption of an item, per minute",
	}, []string{
		"item_name",
	})

	ItemConsumptionCapacityPercent = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "item_consumption_capacity_pc",
		Help: "The percentage of an item's consumption capacity being used",
	}, []string{
		"item_name",
	})

	ItemsProducedPerMin = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "items_produced_per_min",
		Help: "The number of an item being produced, per minute",
	}, []string{
		"item_name",
	})

	ItemsConsumedPerMin = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "items_consumed_per_min",
		Help: "The number of an item being consumed, per minute",
	}, []string{
		"item_name",
	})

	PowerConsumed = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "power_consumed",
		Help: "Power consumed on selected power circuit",
	}, []string{
		"circuit_id",
	})

	PowerCapacity = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "power_capacity",
		Help: "Power capacity on selected power circuit",
	}, []string{
		"circuit_id",
	})

	PowerMaxConsumed = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "power_max_consumed",
		Help: "Maximum Power that can be consumed on selected power circuit",
	}, []string{
		"circuit_id",
	})

	BatteryDifferential = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "battery_differential",
		Help: "Amount of power in excess/deficit going into or out of the battery bank(s). Positive = Charges batteries, Negative = Drains batteries",
	}, []string{
		"circuit_id",
	})

	BatteryPercent = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "battery_percent",
		Help: "Percentage of battery bank(s) charge",
	}, []string{
		"circuit_id",
	})

	BatteryCapacity = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "battery_capacity",
		Help: "Total capacity of battery bank(s)",
	}, []string{
		"circuit_id",
	})

	BatterySecondsEmpty = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "battery_seconds_empty",
		Help: "Seconds until Batteries are empty",
	}, []string{
		"circuit_id",
	})

	BatterySecondsFull = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "battery_seconds_full",
		Help: "Seconds until Batteries are full",
	}, []string{
		"circuit_id",
	})

	FuseTriggered = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "fuse_triggered",
		Help: "Has the fuse been triggered",
	}, []string{
		"circuit_id",
	})
}
