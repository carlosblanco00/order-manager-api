package router

import (
	"github.com/carlosblanco00/order-manager-api/internal/infra/entry-point/api/handler"
	"github.com/labstack/echo/v4"
)

func SetupOrderRoutes(e *echo.Echo, handler handler.OrderHandler) {
	orders := e.Group("/orders")
	orders.POST("/", handler.CreateOrder)
	orders.GET("/:id", handler.GetOrderById)
}
