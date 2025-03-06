package handler

import (
	"net/http"
	"strconv"

	"github.com/carlosblanco00/order-manager-api/internal/domain/model"
	"github.com/carlosblanco00/order-manager-api/internal/domain/usecase"
	"github.com/carlosblanco00/order-manager-api/internal/infra/entry-point/api/helper"
	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	OrderUsecase *usecase.OrderUseCase
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {

	idempotencyKey := c.Request().Header.Get("Idempotency-Key")
	if idempotencyKey == "" {
		return helper.RespondError(c, model.ErrHeaderIdenpotency)
	}

	order := new(model.Order)
	if err := c.Bind(order); err != nil {
		return helper.RespondError(c, model.ErrBadRequest)
	}

	order, err := h.OrderUsecase.CreateOrder(order, idempotencyKey)

	if err != nil {
		return helper.RespondError(c, err)
	}

	return c.JSON(http.StatusCreated, order)
}

func (h *OrderHandler) GetOrderById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return helper.RespondError(c, model.ErrIncorrectID)
	}
	product, err := h.OrderUsecase.GetOrderById(int(id))
	if err != nil {
		return helper.RespondError(c, model.ErrNotFound)
	}
	return c.JSON(http.StatusOK, product)
}
