package main

import (
	"log"

	"github.com/apextrade/config"
	"github.com/apextrade/internal/handlers"
	"github.com/apextrade/internal/repository"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Error while loading env:", err)
	}
	db := config.ConnectDB()

	orderRepo := repository.NewPostgresOrderRepo(db)
	orderHandler := handlers.NewOrderHandler(orderRepo)

	repo := repository.NewInMemoryStockRepo()
	stockHandler := handlers.NewStockHandler(repo)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	apiV1 := r.Group("/api/v1")

	// Order routes
	orderHandler.RegisterRoutes(apiV1)

	stocks := r.Group("/stocks")
	stocks.GET("", stockHandler.GetAllStocks)
	stocks.GET("/:symbol", stockHandler.GetStock)
	stocks.POST("", stockHandler.CreateStock)

	log.Printf("ApexTrade server on port %s – Trade smart!", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
