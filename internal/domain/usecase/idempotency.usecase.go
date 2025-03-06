package usecase

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/carlosblanco00/order-manager-api/internal/domain/model"
)

type IdempotencyUseCase struct {
	RedisGatewway model.RedisGateway
}

func (s *IdempotencyUseCase) EnsureIdempotencyKey(idempotencyKey string) (*model.RedisModel, error) {
	if value, err := s.getExistingKey(idempotencyKey); err != nil || value != nil {
		return value, err
	}

	return s.createInProgressKey(idempotencyKey)
}

func (s *IdempotencyUseCase) getExistingKey(idempotencyKey string) (*model.RedisModel, error) {
	exists, err := s.RedisGatewway.CheckKey(idempotencyKey)
	if err != nil {
		return nil, fmt.Errorf("error checking idempotency key: %w", err)
	}

	if !exists {
		return nil, nil
	}

	payload, err := s.RedisGatewway.GetKey(idempotencyKey)
	if err != nil {
		return nil, model.ManageError(model.ErrNotFound)
	}

	var value model.RedisModel
	if err := json.Unmarshal([]byte(payload), &value); err != nil {
		return nil, fmt.Errorf("error unmarshalling idempotency value: %w", err)
	}

	if value.Status == "IN_PROGRESS" {
		return nil, model.ManageError(model.ErrIdenpotency)
	}

	return &value, nil
}

func (s *IdempotencyUseCase) createInProgressKey(idempotencyKey string) (*model.RedisModel, error) {
	value := model.RedisModel{
		IdempotencyKey: idempotencyKey,
		Status:         "IN_PROGRESS",
		Response:       "",
	}

	if err := s.RedisGatewway.SetKey(idempotencyKey, value, 10*time.Minute); err != nil {
		return nil, fmt.Errorf("error setting idempotency key: %w", err)
	}

	return nil, nil
}

func (s *IdempotencyUseCase) SuccessfulProcess(idempotencyKey string, response string) {
	value := model.RedisModel{
		IdempotencyKey: idempotencyKey,
		Status:         "COMPLETE",
		Response:       response,
	}
	s.RedisGatewway.SetKey(idempotencyKey, value, 90*time.Minute)

}
