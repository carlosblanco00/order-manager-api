package router

import (
	"github.com/carlosblanco00/order-manager-api/internal/infra/entry-point/api/handler"
	"github.com/labstack/echo/v4"
)

func SetupProductRoutes(e *echo.Echo, handler handler.ProductHandler) {
	products := e.Group("/products")
	products.POST("", handler.CreateProduct)
	products.GET("/:id", handler.FindProductById)
	products.GET("", handler.GetAllProducts)
	products.PUT("/:id/stock", handler.UpdateStockProduc)
}
