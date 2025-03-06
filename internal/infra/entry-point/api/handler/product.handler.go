package handler

import (
	"net/http"
	"strconv"

	"github.com/carlosblanco00/order-manager-api/internal/domain/model"
	"github.com/carlosblanco00/order-manager-api/internal/domain/usecase"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	ProductUseCase *usecase.ProductUseCase
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	product := new(model.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product"})
	}

	product, err := h.ProductUseCase.CreateProduct(product)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) FindProductById(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}
	product, err := h.ProductUseCase.GetProductByID(int(id))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Order not found"})
	}
	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) GetAllProducts(c echo.Context) error {

	order, err := h.ProductUseCase.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Order not found"})
	}
	return c.JSON(http.StatusOK, order)
}

func (h *ProductHandler) UpdateStockProduc(c echo.Context) error {

	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid ID format"})
	}

	product := new(model.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid product"})
	}

	product.ID = int(id)

	productUppdate := h.ProductUseCase.UpdateStock(*product)
	return c.JSON(http.StatusOK, productUppdate)
}
