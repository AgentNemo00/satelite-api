package api

import (
	"context"
	router2 "github.com/AgentNemo00/sca-instruments/api/router"
	"github.com/AgentNemo00/sca-instruments/api/router/routen"
	"github.com/AgentNemo00/sca-instruments/api/validation"
	"net/http"
)

//	@title			Sat Image API
//	@version		1.0
//	@description	API for converting geo data to image

//	@host		localhost:10001
//	@BasePath	/

// Start - start api
func StartWithInstruments(ctx context.Context, handler Handler) error {
	validator := validation.NewValidator(func() interface{} {
		return &Payload{}
	})
	base := router2.Simple(
		routen.Basic{
			Method:  http.MethodGet,
			Path:    "/info",
			Handler: handler.Information,
		},
		routen.Validation{
			Basic: routen.Basic{
				Method: http.MethodPost,
				Path:   "/sat",
			},
			Validator: &validator,
			Handler:   handler.Sat,
		},
	)
	apiHandler := router2.NewHandler(&router2.Config{
		Metrics:               true,
		Ping:                  true,
		LoggerSkipInfraRoutes: true,
		Development:           false,
		Port:                  10001,
	})
	err := apiHandler.Build(base)
	if err != nil {
		return err
	}
	return apiHandler.Start()
}
