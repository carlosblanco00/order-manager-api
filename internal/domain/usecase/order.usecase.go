package usecase

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"github.com/carlosblanco00/order-manager-api/internal/domain/model"
)

type OrderUseCase struct {
	OrderGateway       model.OrderGateway
	ProductUseCase     ProductUseCase
	IdempotencyUseCase IdempotencyUseCase
}

func (u *OrderUseCase) CreateOrder(order *model.Order, idempotencyKey string) (*model.Order, error) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var total float64
	var errChan = make(chan error, len(order.OrderItems))
	rollbackProduct := make(chan *model.Product, len(order.OrderItems))

	info, err := u.IdempotencyUseCase.EnsureIdempotencyKey(idempotencyKey)

	if err != nil {
		return nil, err
	}

	if info != nil {
		var existingOrder model.Order
		if err := json.Unmarshal([]byte(info.Response), &existingOrder); err != nil {
			return nil, fmt.Errorf("error unmarshalling idempotency value: %w", err)
		}

		return &existingOrder, nil

	}

	for i := range order.OrderItems {
		wg.Add(1)
		go func(item *model.OrderItem) {
			defer wg.Done()
			err := u.ProductUseCase.HandleStockOperations(item, rollbackProduct)
			if err != nil {
				errChan <- err
				return
			}

			mu.Lock()
			total += item.Subtotal
			mu.Unlock()
		}(&order.OrderItems[i])
	}

	wg.Wait()
	close(errChan)
	close(rollbackProduct)

	for err := range errChan {
		if err != nil {
			for rollback := range rollbackProduct {
				u.ProductUseCase.UpdateStock(*rollback)
			}
			return nil, err
		}
	}

	log.Printf("Order items before save: %+v", order.OrderItems)

	order.TotalAmount = total

	createdOrder, err := u.OrderGateway.Create(order)
	if err != nil {
		return nil, err
	}

	orderJSON, err := json.Marshal(createdOrder)
	if err != nil {
		return nil, fmt.Errorf("error serializing order: %v", err)
	}

	u.IdempotencyUseCase.SuccessfulProcess(idempotencyKey, string(orderJSON))

	return createdOrder, nil
}

func (u *OrderUseCase) GetOrderById(id int) (*model.Order, error) {
	return u.OrderGateway.GetById(id)
}
