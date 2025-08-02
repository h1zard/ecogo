package routes

import (
	"shopping-cart/backend/controllers"
	"shopping-cart/backend/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter configures the routes for the application.
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Add CORS middleware to allow cross-origin requests from the frontend.
	// Make sure the AllowOrigins matches your frontend's URL (e.g., http://localhost:5173 for Vite).
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://deep-shopping-cart.vercel.app"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Public routes
	r.POST("/users", controllers.CreateUser)
	r.POST("/users/login", controllers.Login)
	r.GET("/users", controllers.GetUsers) // Typically, this should be a protected route
	r.GET("/items", controllers.GetItems)
	r.POST("/items", controllers.CreateItem) // Typically, this should be a protected route

	// Protected routes that require authentication
	authRequired := r.Group("/")
	authRequired.Use(middlewares.AuthMiddleware())
	{
		authRequired.POST("/carts", controllers.CreateCart)
		authRequired.GET("/carts", controllers.GetCarts)
		authRequired.POST("/orders", controllers.CreateOrder)
		authRequired.GET("/orders", controllers.GetOrders)
		authRequired.POST("/logout", controllers.Logout)

	}

	return r
}
