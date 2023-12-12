package exporter

import (
	"context"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type PrometheusExporter struct {
	server          *http.Server
	ctx             context.Context
	cancel          context.CancelFunc
	collectorRunner *CollectorRunner
}

func NewPrometheusExporter(frmApiHost string) *PrometheusExporter {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	server := &http.Server{
		Handler: mux,
		Addr:    ":9000",
	}

	ctx, cancel := context.WithCancel(context.Background())

	reg := prometheus.NewRegistry()
	RegisterMetrics(reg)
	RegisterDroneMetrics(reg)
	RegisterFactoryMetrics(reg)
	RegisterTrainMetrics(reg)
	RegisterVehicleMetrics(reg)

	productionCollector := NewProductionCollector(ctx, frmApiHost+"/getProdStats")
	powerCollector := NewPowerCollector(ctx, frmApiHost+"/getPower")
	buildingCollector := NewFactoryBuildingCollector(ctx, frmApiHost+"/getFactory")
	vehicleCollector := NewVehicleCollector(ctx, frmApiHost+"/getVehicles")
	trainCollector := NewTrainCollector(ctx, frmApiHost+"/getTrains")
	droneCollector := NewDroneStationCollector(ctx, frmApiHost+"/getDroneStation")

	collectorRunner := NewCollectorRunner(
		ctx,
		reg,
		productionCollector,
		powerCollector,
		buildingCollector,
		vehicleCollector,
		trainCollector,
		droneCollector,
	)

	return &PrometheusExporter{
		server:          server,
		ctx:             ctx,
		cancel:          cancel,
		collectorRunner: collectorRunner,
	}
}

func (e *PrometheusExporter) Start() {
	go e.collectorRunner.Start()
	go func() {
		err := e.server.ListenAndServe()
		if err != nil {
			log.Print(err)
		}

		log.Println("stopping exporter")
	}()
}

func (e *PrometheusExporter) Stop() error {
	e.cancel()
	return e.server.Close()
}
