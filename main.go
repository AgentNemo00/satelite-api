package main

import (
	"context"
	"github.com/AgentNemo00/satelite-api/api"
	"github.com/AgentNemo00/satelite-api/config"
	"github.com/AgentNemo00/sca-instruments/containerization"
	"log"
	_ "net/http/pprof"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go containerization.Interrupt(cancel)
	cvt, err := config.ConverterByConfig()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting service with config %#v\n", cvt)
	err = api.Start(ctx, api.V1{Converter: cvt})
	if err != nil {
		log.Println(err)
	}
}
