package api

import (
	"context"
	router2 "github.com/AgentNemo00/sca-instruments/api/router"
	"github.com/AgentNemo00/sca-instruments/api/router/routen"
	"github.com/AgentNemo00/sca-instruments/api/validation"
	"github.com/AgentNemo00/sca-instruments/openapi"
	"net/http"
)

type API struct {
	Specs   *openapi.Openapi
	Handler Handler
	Port    int
}

func (a API) Start(ctx context.Context) error {
	validator := validation.NewValidator(func() validation.Model {
		return &Payload{}
	})
	base := router2.SimpleWithCommonPath(
		"",
		a.Specs,
		routen.Basic{
			Method:  http.MethodGet,
			Path:    "/info",
			Handler: a.Handler.Information,
		},
		&routen.Validation{
			Basic: routen.Basic{
				Method: http.MethodPost,
				Path:   "/sat",
			},
			Validator: validator,
			Handler:   a.Handler.Sat,
		},
	)
	apiHandler := router2.NewHandler(&router2.Config{
		Metrics:               true,
		Ping:                  true,
		LoggerSkipInfraRoutes: true,
		Development:           false,
		Port:                  a.Port,
	})
	err := apiHandler.Build(base)
	if err != nil {
		return err
	}
	go func() {
		<-ctx.Done()
		apiHandler.Stop(ctx)
	}()
	return apiHandler.Start()
}
