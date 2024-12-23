package routes

import (
	controllers "github.com/TrizzlyBare/Restaurant-Ordering-System/controllers"
	gin "github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine) {
	r.POST("/payment", controllers.CreatePayment)
	r.GET("/payment/:orderId", controllers.GetPayments)
}