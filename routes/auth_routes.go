package routes

import (
	controllers "github.com/TrizzlyBare/Restaurant-Ordering-System/controllers"
	gin "github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
    r.GET("/user/:id", controllers.GetUser)
    r.PUT("/user/:id", controllers.UpdateUser)
    r.DELETE("/user/:id", controllers.DeleteUser)
}