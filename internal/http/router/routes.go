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

func PublicRoutes(authHandler *handler.AuthHandler) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: authHandler.Login,
		},
	}
}

func PrivateRoutes(userHandler *handler.UserHandler) []*Route {
	return []*Route{
		{
			Method:  echo.GET,
			Path:    "/users",
			Handler: userHandler.GetAllUsers,
		},
		{
			Method:  echo.GET,
			Path:    "/users/:id",
			Handler: userHandler.GetUserByID,
		},
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: userHandler.CreateUser,
		},
		{
			Method:  echo.PUT,
			Path:    "/users/:id",
			Handler: userHandler.UpdateUser,
		},
		{
			Method:  echo.DELETE,
			Path:    "/users/:id",
			Handler: userHandler.DeleteUser,
		},
	}
}
