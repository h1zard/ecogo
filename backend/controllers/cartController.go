package controllers

import (
	"net/http"
	"shopping-cart/backend/database"
	"shopping-cart/backend/models"

	"github.com/gin-gonic/gin"
)

func CreateCart(c *gin.Context) {
	userI, _ := c.Get("user")
	user := userI.(models.User)

	// Read item_id from JSON body
	var input struct {
		ItemID uint `json:"item_id"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var cart models.Cart
	// Check if user already has a cart
	err := database.DB.Where("user_id = ?", user.ID).First(&cart).Error
	if err != nil {
		// If no cart, create one
		cart = models.Cart{UserID: user.ID}
		if err := database.DB.Create(&cart).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create cart"})
			return
		}
	}

	// Add item to cart using CartItem model
	cartItem := models.CartItem{
		CartID: cart.ID,
		ItemID: input.ItemID,
	}
	if err := database.DB.Create(&cartItem).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add item to cart"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Item added to cart", "cart_id": cart.ID})
}

func GetCarts(c *gin.Context) {
	userInterface, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	user := userInterface.(models.User)

	var cart models.Cart
	err := database.DB.Preload("Items.Item").Where("user_id = ?", user.ID).First(&cart).Error
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cart not found"})
		return
	}

	c.JSON(http.StatusOK, cart)
}
