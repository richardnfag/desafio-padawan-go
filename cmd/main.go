package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/richardnfag/desafio-padawan-go/internal/adapters/database"
	"github.com/richardnfag/desafio-padawan-go/internal/adapters/http"
	"github.com/richardnfag/desafio-padawan-go/internal/services"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db := connectDatabase()
	conversionRepo := database.NewGormConversionRepository(db)
	currencyRepo := database.NewGormCurrencyRepository(db)

	conversionService := services.NewConversionService(conversionRepo, currencyRepo)
	httpHandler := http.NewHTTPHandler(conversionService, conversionRepo)

	e := echo.New()
	e.GET("/exchange/:amount/:from/:to/:rate", httpHandler.ConvertCurrency)
	e.Start(":8000")
}

func connectDatabase() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	databaseName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, host, port, databaseName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}
