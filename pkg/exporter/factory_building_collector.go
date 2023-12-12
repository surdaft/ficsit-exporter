package exporter

import (
	"context"
	"strconv"
)

type FactoryBuildingCollector struct {
	FRMAddress string
	ctx        context.Context
}

func NewFactoryBuildingCollector(ctx context.Context, frmAddress string) *FactoryBuildingCollector {
	return &FactoryBuildingCollector{
		FRMAddress: frmAddress,
		ctx:        ctx,
	}
}

func (c *FactoryBuildingCollector) Collect() error {
	details := []BuildingDetail{}
	err := retrieveData(c.ctx, c.FRMAddress, &details)
	if err != nil {
		return err
	}

	for _, building := range details {
		for _, prod := range building.Production {
			MachineItemsProducedPerMin.WithLabelValues(
				prod.Name,
				building.Building,
				strconv.FormatFloat(building.Location.X, 'f', -1, 64),
				strconv.FormatFloat(building.Location.Y, 'f', -1, 64),
				strconv.FormatFloat(building.Location.Z, 'f', -1, 64),
			).Set(prod.CurrentProd)

			MachineItemsProducedEffiency.WithLabelValues(
				prod.Name,
				building.Building,
				strconv.FormatFloat(building.Location.X, 'f', -1, 64),
				strconv.FormatFloat(building.Location.Y, 'f', -1, 64),
				strconv.FormatFloat(building.Location.Z, 'f', -1, 64),
			).Set(prod.ProdPercent)
		}
	}

	return nil
}
