package converter

import (
	"context"
	"github.com/AgentNemo00/satelite-api/model"
)

type SentinelSat struct{}

func (s SentinelSat) Convert(ctx context.Context, data []model.GeoPoint, configuration Configurations) (Result, error) {
	return Result{}, nil
}

func (s SentinelSat) Information(ctx context.Context) Information {
	return Information{
		Name:     "SentinalSat",
		Clipping: true,
		MapTypes: nil,
	}
}
