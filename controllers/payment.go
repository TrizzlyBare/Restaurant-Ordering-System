package controllers

import (
	"net/http"
	"strconv"

	"github.com/TrizzlyBare/Restaurant-Ordering-System/models"
	"github.com/gin-gonic/gin"
)

// CreatePayment handles creating a new payment
func CreatePayment(c *gin.Context) {
    var payment models.Payment
    if err := c.ShouldBindJSON(&payment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := models.DB.Create(&payment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create payment"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Payment created successfully"})
}

// GetPayments handles fetching payments by order ID
func GetPayments(c *gin.Context) {
    orderID, err := strconv.Atoi(c.Param("orderId"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
        return
    }

    var payments []models.Payment
    if err := models.DB.Where("order_id = ?", orderID).Find(&payments).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payments"})
        return
    }

    c.JSON(http.StatusOK, payments)
}