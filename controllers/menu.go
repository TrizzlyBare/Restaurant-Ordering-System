package controllers

import (
	"net/http"
	"strconv"

	"github.com/TrizzlyBare/Restaurant-Ordering-System/models"
	"github.com/gin-gonic/gin"
)

// CreateMenu handles creating a new menu item
func CreateMenu(c *gin.Context) {
    var menu models.Menu
    if err := c.ShouldBindJSON(&menu); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := models.DB.Create(&menu).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create menu item"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Menu item created successfully"})
}

// GetMenus handles fetching all menu items
func GetMenus(c *gin.Context) {
    var menus []models.Menu
    if err := models.DB.Find(&menus).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menu items"})
        return
    }

    c.JSON(http.StatusOK, menus)
}

// GetMenu handles fetching a single menu item by ID
func GetMenu(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
        return
    }

    var menu models.Menu
    if err := models.DB.First(&menu, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Menu item not found"})
        return
    }

    c.JSON(http.StatusOK, menu)
}

// UpdateMenu handles updating a menu item by ID
func UpdateMenu(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
        return
    }

    var menu models.Menu
    if err := models.DB.First(&menu, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Menu item not found"})
        return
    }

    if err := c.ShouldBindJSON(&menu); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := models.DB.Save(&menu).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update menu item"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Menu item updated successfully"})
}

// DeleteMenu handles deleting a menu item by ID
func DeleteMenu(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid menu ID"})
        return
    }

    if err := models.DB.Delete(&models.Menu{}, id).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete menu item"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Menu item deleted successfully"})
}