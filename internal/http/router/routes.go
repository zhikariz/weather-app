package router

import (
	"github.com/labstack/echo/v4"
	"github.com/zhikariz/weather-app/internal/http/handler"
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

func PublicRoutes(userHandler *handler.UserHandler) []*Route {
	return []*Route{
		{
			Method:  echo.GET,
			Path:    "/users",
			Handler: userHandler.GetAllUsers,
		},
	}
}

func PrivateRoutes() []*Route {
	return []*Route{}
}
