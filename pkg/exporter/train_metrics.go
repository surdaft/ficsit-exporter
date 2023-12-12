package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	TrainRoundTrip   *prometheus.GaugeVec
	TrainSegmentTrip *prometheus.GaugeVec
	TrainDerailed    *prometheus.GaugeVec
	TrainPower       *prometheus.GaugeVec
)

func RegisterTrainMetrics(reg *prometheus.Registry) {
	TrainRoundTrip = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "train_round_trip_seconds",
		Help: "Recorded train round trip time in seconds",
	}, []string{
		"name",
	})
	TrainSegmentTrip = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "train_segment_trip_seconds",
		Help: "Recorded train trip between two stations",
	}, []string{
		"name",
		"from",
		"to",
	})
	TrainDerailed = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "train_derailed",
		Help: "Is train derailed",
	}, []string{
		"name",
	})
	TrainPower = RegisterNewGaugeVec(reg, prometheus.GaugeOpts{
		Name: "train_power_consumed",
		Help: "How much power train is consuming",
	}, []string{
		"name",
	})
}
