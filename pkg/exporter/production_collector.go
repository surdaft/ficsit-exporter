package exporter

import "context"

type ProductionCollector struct {
	FRMAddress string
	ctx        context.Context
}

type ProductionDetails struct {
	ItemName           string  `json:"ItemName"`
	ProdPercent        float64 `json:"ProdPercent"`
	ConsPercent        float64 `json:"ConsPercent"`
	CurrentProduction  float64 `json:"CurrentProd"`
	CurrentConsumption float64 `json:"CurrentConsumed"`
	MaxProd            float64 `json:"MaxProd"`
	MaxConsumed        float64 `json:"MaxConsumed"`
}

func NewProductionCollector(ctx context.Context, frmAddress string) *ProductionCollector {
	return &ProductionCollector{
		FRMAddress: frmAddress,
		ctx:        ctx,
	}
}

func (c *ProductionCollector) Collect() error {
	details := []ProductionDetails{}
	err := retrieveData(c.ctx, c.FRMAddress, &details)
	if err != nil {
		return err
	}

	for _, d := range details {
		ItemsProducedPerMin.WithLabelValues(d.ItemName).Set(d.CurrentProduction)
		ItemsConsumedPerMin.WithLabelValues(d.ItemName).Set(d.CurrentConsumption)

		ItemProductionCapacityPercent.WithLabelValues(d.ItemName).Set(d.ProdPercent)
		ItemConsumptionCapacityPercent.WithLabelValues(d.ItemName).Set(d.ConsPercent)
		ItemProductionCapacityPerMinute.WithLabelValues(d.ItemName).Set(d.MaxProd)
		ItemConsumptionCapacityPerMinute.WithLabelValues(d.ItemName).Set(d.MaxConsumed)
	}

	return nil
}
