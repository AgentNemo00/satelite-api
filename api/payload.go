package api

import (
	"encoding/json"
	"github.com/AgentNemo00/satelite-api/converter"
	"github.com/AgentNemo00/satelite-api/model"
)

type Payload struct {
	Coordinates []model.GeoPoint `json:"coordinates"`
	converter.Configurations
}

func (p *Payload) Parse(data []byte) error {
	return json.Unmarshal(data, p)
}
