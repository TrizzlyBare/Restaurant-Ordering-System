package controllers

import (
	"net/http"
	"strconv"

	"github.com/TrizzlyBare/Restaurant-Ordering-System/models"
	"github.com/gin-gonic/gin"
)

// CreateOrder handles creating a new order
func CreateOrder(c *gin.Context) {
    var order models.Order
    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := models.DB.Create(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order created successfully"})
}

// GetOrders handles fetching all orders
func GetOrders(c *gin.Context) {
    var orders []models.Order
    if err := models.DB.Find(&orders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
        return
    }

    c.JSON(http.StatusOK, orders)
}

// GetOrder handles fetching a single order by ID
func GetOrder(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
        return
    }

    var order models.Order
    if err := models.DB.First(&order, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    c.JSON(http.StatusOK, order)
}

// UpdateOrder handles updating an order by ID
func UpdateOrder(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
        return
    }

    var order models.Order
    if err := models.DB.First(&order, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
    }

    if err := c.ShouldBindJSON(&order); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := models.DB.Save(&order).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}

// DeleteOrder handles deleting an order by ID
func DeleteOrder(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
        return
    }

    if err := models.DB.Delete(&models.Order{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete order"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}