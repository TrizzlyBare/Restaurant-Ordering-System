package routes

import (
	controllers "github.com/TrizzlyBare/Restaurant-Ordering-System/controllers"
	gin "github.com/gin-gonic/gin"
)

func MenuRoutes(r *gin.Engine) {
	r.POST("/menu", controllers.CreateMenu)
	r.GET("/menu", controllers.GetMenus)
	r.GET("/menu/:id", controllers.GetMenu)
	r.PUT("/menu/:id", controllers.UpdateMenu)
	r.DELETE("/menu/:id", controllers.DeleteMenu)
}