package controllers

import (
	"net/http"
	"shopping-cart/backend/database"
	"shopping-cart/backend/models"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	var items []models.Item

	if err := c.ShouldBindJSON(&items); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := database.DB.Create(&items); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

func GetItems(c *gin.Context) {
	var items []models.Item
	database.DB.Find(&items)
	c.JSON(http.StatusOK, items)
}
