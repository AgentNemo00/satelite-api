package api

import (
	"github.com/AgentNemo00/satelite-api/converter"
	"github.com/AgentNemo00/sca-instruments/log"
	"github.com/kataras/iris/v12/context"
	"net/http"
)

type Handler struct {
	Converter converter.Converter
}

// Info godoc
//
//	@Summary		Information
//	@Description	Informations and Feature set of the converter
//	@Tags			V1
//	@Produce		json
//	@Success		200 {object} converter.Information
//	@Failure		500
//	@Router			/v1/info [get]
func (h Handler) Information(c *context.Context) {
	err := c.StopWithJSON(http.StatusOK, h.Converter.Information(c))
	if err != nil {
		log.Ctx(c).Error(err.Error())
	}
}

// Convert godoc
//
//	@Summary		Generate Image
//	@Description	Generates an image from geo data
//	@Tags			V1
//	@Accept			json
//	@Produce		png, application/zip
//	@Param			data body Payload true " "
//	@Success		200
//	@Failure		409
//	@Failure		500
//	@Router			/v1/sat [post]
func (h Handler) Sat(c *context.Context, data interface{}) {
	payload := data.(*Payload)

	if len(payload.Coordinates) < 3 {
		c.StopWithStatus(http.StatusConflict)
		return
	}
	result, err := h.Converter.Convert(c, payload.Coordinates, payload.Configurations)
	if err != nil {
		c.StopWithStatus(http.StatusInternalServerError)
		log.Ctx(c).Error(err.Error())
		return
	}
	err = result.ParseResponse(c, c.ResponseWriter())
	if err != nil {
		c.StopWithStatus(http.StatusInternalServerError)
		log.Ctx(c).Error(err.Error())
		return
	}
	c.StopWithStatus(http.StatusOK)
}
