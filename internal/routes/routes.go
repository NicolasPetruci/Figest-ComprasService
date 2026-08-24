package routes

import (
	"github.com/NicolasPetruci/Figest-ComprasService/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	r.GET("/health", handlers.HealthCheck)
	
	purchases := r.Group("/purchases")
	{
		purchases.POST("", handlers.CreatePurchase)
		purchases.GET("", handlers.GetPurchases)
		purchases.GET("/summary", handlers.GetPurchasesSummary)
		purchases.GET("/:id", handlers.GetPurchase)
		purchases.PATCH("/:id", handlers.UpdatePurchase)
		purchases.DELETE("/:id", handlers.DeletePurchase)
	}

	suppliers := r.Group("/suppliers")
	{
		suppliers.POST("", handlers.CreateSupplier)
		suppliers.GET("", handlers.GetSuppliers)
		suppliers.GET("/:id", handlers.GetSupplier)
		suppliers.GET("/:id/purchases", handlers.GetSupplierPurchases)
	}
}
