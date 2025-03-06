package main

import (
	"github.com/carlosblanco00/order-manager-api/internal/domain/usecase"
	"github.com/carlosblanco00/order-manager-api/internal/infra/driven-adapter/mysql/config"
	"github.com/carlosblanco00/order-manager-api/internal/infra/driven-adapter/mysql/repository"
	"github.com/carlosblanco00/order-manager-api/internal/infra/driven-adapter/redis"
	redisRepo "github.com/carlosblanco00/order-manager-api/internal/infra/driven-adapter/redis/repository"
	"github.com/carlosblanco00/order-manager-api/internal/infra/entry-point/api/handler"
	"github.com/carlosblanco00/order-manager-api/internal/infra/entry-point/api/router"
	"github.com/labstack/echo/v4"
)

func main() {

	config.InitDB()

	redis.InitRedisConnection()

	e := echo.New()

	orderRepo := &repository.OrderRepository{
		Db: config.DB,
	}

	productRepo := &repository.ProductRepositpry{
		Db: config.DB,
	}

	redisRepo := &redisRepo.RedisRepository{
		Client: redis.Client,
	}

	productUseCase := &usecase.ProductUseCase{
		ProductGateway: productRepo,
	}

	idempotencyUseCase := &usecase.IdempotencyUseCase{
		RedisGatewway: redisRepo,
	}

	OrderUseCase := &usecase.OrderUseCase{
		OrderGateway:       orderRepo,
		ProductUseCase:     *productUseCase,
		IdempotencyUseCase: *idempotencyUseCase,
	}

	orderHandler := &handler.OrderHandler{
		OrderUsecase: OrderUseCase,
	}

	productHandler := &handler.ProductHandler{
		ProductUseCase: productUseCase,
	}

	router.SetupOrderRoutes(e, *orderHandler)
	router.SetupProductRoutes(e, *productHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
