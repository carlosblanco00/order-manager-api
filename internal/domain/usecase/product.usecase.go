package usecase

import (
	"fmt"
	"log"

	"github.com/carlosblanco00/order-manager-api/internal/domain/model"
)

type ProductUseCase struct {
	ProductGateway model.ProductGateway
}

func (p *ProductUseCase) GetProductByID(id int) (*model.Product, error) {
	return p.ProductGateway.GetById(id)
}

func (p *ProductUseCase) GetAllProducts() ([]*model.Product, error) {
	return p.ProductGateway.GetAll()
}

func (p *ProductUseCase) UpdateStock(product model.Product) *model.Product {

	productUpdate, err := p.FindProductById(product.ID)
	log.Printf("erro: %v product: %v", err, productUpdate)
	if err == nil {
		productUpdate.Stock = product.Stock
	}

	p.ProductGateway.Update(productUpdate)
	return productUpdate
}

func (p *ProductUseCase) CreateProduct(product *model.Product) (*model.Product, error) {
	return p.ProductGateway.Create(product)
}

func (p *ProductUseCase) HandleStockOperations(item *model.OrderItem, rollbackProduct chan *model.Product) error {
	product, err := p.GetProductByID(item.ProductID)
	if err != nil {
		return model.ManageError(model.ErrNotFound)
	}

	if product.Stock < item.Quantity {
		return model.ManageError(model.ErrStock)

	}

	rollbackProduct <- &model.Product{ID: product.ID, Stock: product.Stock}
	product.Stock -= item.Quantity
	if err := p.ProductGateway.Update(product); err != nil {
		return fmt.Errorf("failed to update stock: %w", err)
	}

	item.Subtotal = float64(item.Quantity) * product.Price

	log.Printf("product price: %v", product.Price)
	log.Printf("subtotal: %v", item.Subtotal)

	return nil
}

func (p *ProductUseCase) FindProductById(id int) (*model.Product, error) {

	product, err := p.GetProductByID(id)
	if err != nil {
		return nil, fmt.Errorf("product not found for ID: %d error: %v", id, err)
	}

	return product, nil
}
