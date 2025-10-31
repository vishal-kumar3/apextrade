package handlers

import (
	"net/http"
	"strconv"

	"github.com/apextrade/internal/models"
	"github.com/apextrade/internal/repository"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	repo *repository.PostgresOrderRepo
}

func NewOrderHandler(repo *repository.PostgresOrderRepo) *OrderHandler {
	return &OrderHandler{repo: repo}
}

func (h *OrderHandler) RegisterRoutes(rg *gin.RouterGroup) {
	orderRouteGroup := rg.Group("/orders")
	{
		orderRouteGroup.POST("/", h.CreateOrderHandler)
		orderRouteGroup.GET("/:id", h.GetOrderByID)
		orderRouteGroup.GET("/stock/:symbol", h.GetOrderByStock)
		orderRouteGroup.PUT("/:id/status", h.UpdateOrderStatus)
	}
}

func (h *OrderHandler) CreateOrderHandler(c *gin.Context) {
	var o models.Order
	if err := c.ShouldBindJSON(&o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.repo.CreateOrder(&o); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, o)
}

func (h *OrderHandler) GetOrderByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	order, exists := h.repo.GetByID(uint(id))
	if !exists {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found!"})
		return
	}
	c.IndentedJSON(http.StatusOK, order)
}

func (h *OrderHandler) GetOrderByStock(c *gin.Context) {
	symbol := c.Param("symbol")
	orders := h.repo.GetBySymbol(symbol)
	c.IndentedJSON(http.StatusOK, orders)
}

func (h *OrderHandler) UpdateOrderStatus(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req struct {
		Status models.OrderStatus `json:"status" validate:"required,oneof=filled cancelled"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if !h.repo.UpdateStatus(uint(id), req.Status) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found!"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Order status updated"})
}
