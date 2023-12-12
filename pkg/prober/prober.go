package prober

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/AP-Hunt/ficsit-exporter/pkg/exporter"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

type Prober struct {
	UserAgent    string
	CobraContext *cobra.Command
}

func New(cmd *cobra.Command) *Prober {
	return &Prober{
		UserAgent:    fmt.Sprintf(cmd.Version),
		CobraContext: cmd,
	}
}

func (p *Prober) Handle(c *gin.Context) {
	// we need a target host and port
	targetAddress := c.Query("address")

	parsedUrl, err := url.Parse("http://" + targetAddress)
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{
			"message": "invalid url: " + err.Error(),
		})
	}

	// create a new registry with every probe, so that the next probe is fresh
	reg := prometheus.NewRegistry()

	// share the request context all the way through. if the request is cancelled
	// then all the probe requests should be too
	p.probe(c.Request.Context(), reg, parsedUrl)
	promhttp.HandlerFor(reg, promhttp.HandlerOpts{}).ServeHTTP(c.Writer, c.Request)
}

func (p *Prober) probe(ctx context.Context, reg *prometheus.Registry, address *url.URL) {
	g := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "frmc_success",
	})

	reg.MustRegister(g)

	exporter.RegisterMetrics(reg)
	exporter.RegisterDroneMetrics(reg)
	exporter.RegisterFactoryMetrics(reg)
	exporter.RegisterTrainMetrics(reg)
	exporter.RegisterVehicleMetrics(reg)

	productionCollector := exporter.NewProductionCollector(ctx, address.String()+"/getProdStats")
	powerCollector := exporter.NewPowerCollector(ctx, address.String()+"/getPower")
	buildingCollector := exporter.NewFactoryBuildingCollector(ctx, address.String()+"/getFactory")
	vehicleCollector := exporter.NewVehicleCollector(ctx, address.String()+"/getVehicles")
	trainCollector := exporter.NewTrainCollector(ctx, address.String()+"/getTrains")
	droneCollector := exporter.NewDroneStationCollector(ctx, address.String()+"/getDroneStation")

	collectorRunner := exporter.NewCollectorRunner(
		ctx,
		reg,
		productionCollector,
		powerCollector,
		buildingCollector,
		vehicleCollector,
		trainCollector,
		droneCollector,
	)

	errs := make(chan error)
	collectorRunner.Collect(errs)
	g.Set(1)

	go func(e chan error, g *prometheus.Gauge) {
		for {
			select {
			case <-e:
				// log.Print(err)
				gg := *g
				gg.Set(0)
			}
		}
	}(errs, &g)
}
