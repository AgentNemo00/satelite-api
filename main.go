package main

import (
	"context"
	"log"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sat-api/api"
	"sat-api/converter"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		<-c
		cancel()
	}()
	err := api.Start(ctx, api.V1{Converter: converter.OpenStreet{
		Name:                    "OpenStreet",
		Signature:               "",
		OptimizedSizeToleration: 10,
		AreaWeight:              1.5,
		DefaultHeight:           512,
		DefaultWidth:            512,
		MaxZoom:                 20,
		MaximalArea:             25340.20196778653, // e-10
		GeoDataMultiplier:       10000000,
		Cache:                   nil, //sm.NewTileCacheFromUserCache(0777),
		ParallelAtNumberOf:      3,
	}})
	if err != nil {
		log.Println(err)
	}
}
