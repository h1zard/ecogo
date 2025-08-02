package controllers

import (
	"net/http"
	"shopping-cart/backend/database"
	"shopping-cart/backend/models"

	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	user, _ := c.Get("user")

	var order models.Order
	if err := c.ShouldBindJSON(&order); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.UserID = user.(models.User).ID
	database.DB.Create(&order)

	c.JSON(http.StatusOK, order)
}

func GetOrders(c *gin.Context) {
	var orders []models.Order
	database.DB.Find(&orders)
	c.JSON(http.StatusOK, orders)
}
