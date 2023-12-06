package builder

import (
	"github.com/go-redis/redis/v8"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/zhikariz/weather-app/internal/config"
	"github.com/zhikariz/weather-app/internal/http/handler"
	"github.com/zhikariz/weather-app/internal/http/router"
	"github.com/zhikariz/weather-app/internal/repository"
	"github.com/zhikariz/weather-app/internal/service"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB, midtransClient snap.Client, redisClient *redis.Client) []*router.Route {
	userRepository := repository.NewUserRepository(db, redisClient)
	transactionRepository := repository.NewTransactionRepository(db)

	loginService := service.NewLoginService(userRepository)
	tokenService := service.NewTokenService(cfg)
	transactionService := service.NewTransactionService(transactionRepository)
	paymentService := service.NewPaymentService(midtransClient)

	authHandler := handler.NewAuthHandler(loginService, tokenService)
	transactionHandler := handler.NewTransactionHandler(transactionService, paymentService)
	return router.PublicRoutes(authHandler, transactionHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB, midtransClient snap.Client, redisClient *redis.Client) []*router.Route {
	userRepository := repository.NewUserRepository(db, redisClient)
	transactionRepository := repository.NewTransactionRepository(db)

	paymentService := service.NewPaymentService(midtransClient)
	transactionService := service.NewTransactionService(transactionRepository)
	userService := service.NewUserService(userRepository)

	userHandler := handler.NewUserHandler(cfg, userService)
	transactionHandler := handler.NewTransactionHandler(transactionService, paymentService)

	return router.PrivateRoutes(userHandler, transactionHandler)
}
