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

func (h Handler) Information(c *context.Context) {
	err := c.StopWithJSON(http.StatusOK, h.Converter.Information(c))
	if err != nil {
		log.Ctx(c).Error(err.Error())
	}
}

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
