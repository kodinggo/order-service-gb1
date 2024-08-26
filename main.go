package main

import (
	"fmt"
	"order-service-gb1/db/db"
	"order-service-gb1/internal/handler"
	"order-service-gb1/internal/repository"
	"order-service-gb1/internal/services"

	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	//init database
	mysqlDB := db.ConfigDB()

	//sClose the connection when done
	sqlDB, err := mysqlDB.DB()
	if err != nil {
		panic(fmt.Errorf("failed to get *sql.DB from *gorm.DB: %w", err))
	}
	defer sqlDB.Close()

	cartRepo := repository.NewCartsRepository(mysqlDB)

	cartService := services.NewCartsRepository(cartRepo)

	e := echo.New()
	e.Use(echoMiddleware.Recover())
	e.Use(echoMiddleware.Logger())

	routeGroup := e.Group("/api/v1")

	handler.NewCartsRepository(routeGroup, cartService)

	e.Logger.Fatal(e.Start(":3200"))
}
