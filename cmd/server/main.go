package main

import (
	"log"

	"github.com/apextrade/config"
	"github.com/apextrade/internal/handlers"
	"github.com/apextrade/internal/models"
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

	stockRepo := repository.NewPostgresStockRepo(db)
	stockHandler := handlers.NewStockHandler(stockRepo)

	seedStocks := []models.Stock{
		{Symbol: "AAPL", Price: 150.25, Volume: 1000000},
		{Symbol: "GOOG", Price: 2800.50, Volume: 500000},
	}
	for _, s := range seedStocks {
		if err := stockRepo.CreateOrUpdate(&s); err != nil {
			log.Printf("Seed warning: %v", err)
		}
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	apiV1 := r.Group("/api/v1")

	// API Routes
	orderHandler.RegisterRoutes(apiV1)
	stockHandler.RegisterRoutes(apiV1)

	log.Printf("ApexTrade server on port %s – Trade smart!", cfg.Server.Port)
	if err := r.Run(":" + cfg.Server.Port); err != nil {
		log.Fatal(err)
	}
}
