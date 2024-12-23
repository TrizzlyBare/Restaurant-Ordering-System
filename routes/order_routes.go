package routes

import (
	controllers "github.com/TrizzlyBare/Restaurant-Ordering-System/controllers"
	gin "github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {
	r.POST("/order", controllers.CreateOrder)
	r.GET("/order", controllers.GetOrders)
	r.GET("/order/:id", controllers.GetOrder)
	r.PUT("/order/:id", controllers.UpdateOrder)
	r.DELETE("/order/:id", controllers.DeleteOrder)
}