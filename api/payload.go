package api

import (
	"github.com/AgentNemo00/satelite-api/converter"
	"github.com/AgentNemo00/satelite-api/model"
)

type Payload struct {
	Coordinates []model.GeoPoint `json:"coordinates"`
	converter.Configurations
}
