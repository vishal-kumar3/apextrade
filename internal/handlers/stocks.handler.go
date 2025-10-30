package handlers

import (
	"net/http"

	"github.com/apextrade/internal/models"
	"github.com/apextrade/internal/repository"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type StockHandler struct {
	repo *repository.InMemoryStockRepo
}

func NewStockHandler(repo *repository.InMemoryStockRepo) *StockHandler {
	return &StockHandler{repo: repo}
}

func (h *StockHandler) GetAllStocks(c *gin.Context) {
	stocks := h.repo.GetAll()
	c.IndentedJSON(http.StatusOK, stocks)
}

func (h *StockHandler) GetStock(c *gin.Context) {
	symbol := c.Param("symbol")
	stock, found := h.repo.GetBySymbol(symbol)
	if !found {
		c.IndentedJSON(http.StatusNotFound, gin.H{"error": "Stock not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, stock)
}

func (h *StockHandler) CreateStock(c *gin.Context) {
	var s models.Stock
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	v := validator.New()
	if err := v.Struct(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	h.repo.CreateOrUpdate(&s)
	c.IndentedJSON(http.StatusCreated, s)
}
