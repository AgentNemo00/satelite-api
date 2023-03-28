package main

import (
	"context"
	"github.com/AgentNemo00/satelite-api/api"
	"github.com/AgentNemo00/satelite-api/assets"
	"github.com/AgentNemo00/satelite-api/config"
	"github.com/AgentNemo00/sca-instruments/configuration"
	"github.com/AgentNemo00/sca-instruments/containerization"
	"github.com/AgentNemo00/sca-instruments/openapi"
	"log"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go containerization.Interrupt(cancel)
	generalConfiguration, err := configuration.ByEnv(&config.Config{})
	if err != nil {
		log.Fatal(err)
	}
	applicationConfiguration := generalConfiguration.(*config.Config)
	cvt, err := config.ConverterByConfig(applicationConfiguration)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting service with config %#v\n", cvt)
	err = api.API{
		Specs: &openapi.Openapi{
			Data: assets.Data(),
			Type: openapi.YAML,
		},
		Handler: api.Handler{Converter: cvt},
		Port:    applicationConfiguration.Port,
	}.Start(ctx)
	if err != nil {
		log.Println(err)
	}
}
