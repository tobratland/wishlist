package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tobratland/wishlist/backend/controllers"
	"github.com/tobratland/wishlist/backend/middleware"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		// Authentication routes
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		// Protected routes
		protected := api.Group("/")
		protected.Use(middleware.AuthMiddleware())
		{
			// Wishlist management
			protected.POST("/wishlists", controllers.CreateWishlist)
			protected.GET("/wishlists/:id", controllers.GetWishlist)
			protected.POST("/wishlists/:id/share", controllers.ShareWishlist)
			protected.POST("/wishlists/:id/items", controllers.AddItem)
			protected.PUT("/items/:id/purchase", controllers.PurchaseItem)
		}

		// Shared wishlist access (public)
		api.GET("/shared/:token", controllers.GetSharedWishlist)
	}
}
