package http

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/richardnfag/desafio-padawan-go/internal/ports"
	"github.com/richardnfag/desafio-padawan-go/internal/services"
)

type HTTPHandler struct {
	conversionService    services.ConversionService
	conversionRepository ports.ConversionRepository
}

func NewHTTPHandler(conversionService services.ConversionService, conversionRepository ports.ConversionRepository) *HTTPHandler {
	return &HTTPHandler{conversionService: conversionService, conversionRepository: conversionRepository}
}

func (h *HTTPHandler) ConvertCurrency(c echo.Context) error {
	amount, _ := strconv.ParseFloat(c.Param("amount"), 64)
	fromCurrency := c.Param("from")
	toCurrency := c.Param("to")
	rate, _ := strconv.ParseFloat(c.Param("rate"), 64)

	conversion, err := h.conversionService.Convert(amount, fromCurrency, toCurrency, rate)

	if err != nil {
		e := map[string]interface{}{
			"error": err.Error(),
		}
		return c.JSON(http.StatusNotFound, e)
	}

	response := map[string]interface{}{
		"valorConvertido": conversion.Result,
		"simboloMoeda":    conversion.To.Symbol,
	}

	return c.JSON(http.StatusOK, response)
}
