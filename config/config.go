package config

import (
	"fmt"
	"github.com/AgentNemo00/satelite-api/converter"
	"github.com/AgentNemo00/sca-instruments/configuration"
	sm "github.com/flopp/go-staticmaps"
)

type Config struct {
	Port                    int
	ConverterName           string
	Signature               string
	OptimizedSizeToleration float64
	AreaWeight              float64
	Height                  int
	Width                   int
	MaxZoom                 int
	MaximalArea             float64
	GeoDatMultiplier        float64
	Cache                   bool
	Parallel                bool
}

func (c *Config) Default() {
	if c.Port == 0 {
		c.Port = 10001
	}
	if c.ConverterName == "" {
		c.ConverterName = "OpenStreet"
	}
	if c.AreaWeight == 0 {
		c.AreaWeight = 1.5
	}
	if c.Height == 0 {
		c.Height = 512
	}
	if c.Width == 0 {
		c.Width = 512
	}
	if c.MaxZoom == 0 {
		c.MaxZoom = 20
	}
	if c.MaximalArea == 0 {
		c.MaximalArea = 25340.20196778653 // e-10
	}
	if c.GeoDatMultiplier == 0 {
		c.GeoDatMultiplier = 10000000
	}
}

func ConverterByConfig(config *Config) (converter.Converter, error) {
	c, err := configuration.ByEnv(&Config{})
	if err != nil {
		return nil, err
	}
	config, ok := c.(*Config)
	if !ok {
		return nil, fmt.Errorf("could not convert config")
	}
	var cvt converter.Converter
	switch config.ConverterName {
	case "OpenStreet":
		var cache sm.TileCache
		if config.Cache {
			cache = sm.NewTileCacheFromUserCache(0777)
		}
		parallelNumberOf := 0
		if config.Parallel {
			parallelNumberOf = 3
		}
		cvt = converter.OpenStreet{
			Name:                    config.ConverterName,
			Signature:               config.Signature,
			OptimizedSizeToleration: config.OptimizedSizeToleration,
			AreaWeight:              config.AreaWeight,
			DefaultHeight:           config.Height,
			DefaultWidth:            config.Width,
			MaxZoom:                 config.MaxZoom,
			MaximalArea:             config.MaximalArea,
			GeoDataMultiplier:       config.GeoDatMultiplier,
			Cache:                   cache,
			ParallelAtNumberOf:      parallelNumberOf,
		}
	default:
		return nil, fmt.Errorf("converter with name %s not found", config.ConverterName)
	}
	return cvt, nil
}
