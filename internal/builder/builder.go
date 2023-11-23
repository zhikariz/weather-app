package builder

import (
	"github.com/zhikariz/weather-app/internal/config"
	"github.com/zhikariz/weather-app/internal/http/handler"
	"github.com/zhikariz/weather-app/internal/http/router"
	"github.com/zhikariz/weather-app/internal/repository"
	"github.com/zhikariz/weather-app/internal/service"
	"gorm.io/gorm"
)

func BuildPublicRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	userRepository := repository.NewUserRepository(db)
	loginService := service.NewLoginService(userRepository)
	tokenService := service.NewTokenService(cfg)
	authHandler := handler.NewAuthHandler(loginService, tokenService)
	return router.PublicRoutes(authHandler)
}

func BuildPrivateRoutes(cfg *config.Config, db *gorm.DB) []*router.Route {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	return router.PrivateRoutes(userHandler)
}
