package http

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/richardnfag/desafio-padawan-go/internal/adapters/database"
	"github.com/richardnfag/desafio-padawan-go/internal/entities"
	"github.com/richardnfag/desafio-padawan-go/internal/services"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var urlPath string = "/exchange/:amount/:from/:to/:rate"

func setupHandler() *HTTPHandler {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&entities.Conversion{}, &entities.Currency{})

	conversionRepo := database.NewGormConversionRepository(db)
	currencyRepo := database.NewGormCurrencyRepository(db)

	insertCurrencies(currencyRepo)

	conversionService := services.NewConversionService(conversionRepo, currencyRepo)
	httpHandler := NewHTTPHandler(conversionService, conversionRepo)

	return httpHandler
}

func insertCurrencies(currencyRepo *database.GormCurrencyRepository) {
	currencies := []entities.Currency{
		{Code: "USD", Symbol: "$"},
		{Code: "BRL", Symbol: "R$"},
		{Code: "EUR", Symbol: "€"},
		{Code: "BTC", Symbol: "₿"},
	}

	for _, currency := range currencies {
		if err := currencyRepo.SaveCurrency(&currency); err != nil {
			log.Println("failed to insert currency: " + err.Error())
		}
	}
}

// Teste De Real para Dólar;
// /exchange/10/BRL/USD/0.20
func TestGetExchangeBRLToUSD(t *testing.T) {
	httpHandler := setupHandler()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(urlPath)
	c.SetParamNames("amount", "from", "to", "rate")
	c.SetParamValues("10", "BRL", "USD", "0.20")

	if assert.NoError(t, httpHandler.ConvertCurrency(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, "{\"valorConvertido\":2,\"simboloMoeda\":\"$\"}\n", rec.Body.String())
	}

}

// * De Dólar para Real;
// /exchange/10/USD/BRL/5.00
func TestGetExchangeUSDToBRL(t *testing.T) {
	httpHandler := setupHandler()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(urlPath)
	c.SetParamNames("amount", "from", "to", "rate")
	c.SetParamValues("10", "USD", "BRL", "5.00")

	if assert.NoError(t, httpHandler.ConvertCurrency(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, "{\"valorConvertido\":50,\"simboloMoeda\":\"R$\"}\n", rec.Body.String())
	}

}

// * De Real para Euro;
// /exchange/10/BRL/EUR/0.18
func TestGetExchangeBRLToEUR(t *testing.T) {
	httpHandler := setupHandler()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(urlPath)
	c.SetParamNames("amount", "from", "to", "rate")
	c.SetParamValues("10", "BRL", "EUR", "0.16")

	if assert.NoError(t, httpHandler.ConvertCurrency(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, "{\"valorConvertido\":1.6,\"simboloMoeda\":\"€\"}\n", rec.Body.String())
	}

}

// * De Euro para Real;
// /exchange/10/EUR/BRL/5.40
func TestGetExchangeEURToBRL(t *testing.T) {
	httpHandler := setupHandler()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(urlPath)
	c.SetParamNames("amount", "from", "to", "rate")
	c.SetParamValues("10", "EUR", "BRL", "5.40")

	if assert.NoError(t, httpHandler.ConvertCurrency(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, "{\"valorConvertido\":54,\"simboloMoeda\":\"R$\"}\n", rec.Body.String())
	}

}

// * De BTC para Dolar;
// /exchange/10/BTC/USD/30000.00
func TestGetExchangeBTCToUSD(t *testing.T) {
	httpHandler := setupHandler()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(urlPath)
	c.SetParamNames("amount", "from", "to", "rate")
	c.SetParamValues("0.05", "BTC", "USD", "30000.00")

	if assert.NoError(t, httpHandler.ConvertCurrency(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, "{\"valorConvertido\":1500,\"simboloMoeda\":\"$\"}\n", rec.Body.String())
	}

}

// * De BTC para Real;
// /exchange/10/BTC/BRL/145000.00
func TestGetExchangeBTCToBRL(t *testing.T) {
	httpHandler := setupHandler()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(urlPath)
	c.SetParamNames("amount", "from", "to", "rate")
	c.SetParamValues("0.05", "BTC", "BRL", "145000.00")

	if assert.NoError(t, httpHandler.ConvertCurrency(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.JSONEq(t, "{\"valorConvertido\":7250,\"simboloMoeda\":\"R$\"}\n", rec.Body.String())
	}

}

// FromCurrency Unknown
func TestGetExchangeFromCurrencyUnknown(t *testing.T) {
	httpHandler := setupHandler()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(urlPath)
	c.SetParamNames("amount", "from", "to", "rate")
	c.SetParamValues("10", "UNK", "BRL", "5.50")

	if assert.NoError(t, httpHandler.ConvertCurrency(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.JSONEq(t, "{\"error\":\"currency code UNK not found\"}\n", rec.Body.String())
	}

}

// ToCurrency Unknown
func TestGetExchangeToCurrencyUnknown(t *testing.T) {
	httpHandler := setupHandler()

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(urlPath)
	c.SetParamNames("amount", "from", "to", "rate")
	c.SetParamValues("10", "BTC", "UNK", "5.50")

	if assert.NoError(t, httpHandler.ConvertCurrency(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.JSONEq(t, "{\"error\":\"currency code UNK not found\"}\n", rec.Body.String())
	}

}
