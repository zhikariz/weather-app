package router

import (
	"github.com/labstack/echo/v4"
	"github.com/zhikariz/weather-app/internal/http/handler"
)

const (
	Administrator = "Administrator"
	Editor        = "Editor"
)

var (
	allRoles   = []string{Administrator, Editor}
	onlyAdmin  = []string{Administrator}
	onlyEditor = []string{Editor}
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
	Roles   []string
}

func PublicRoutes(authHandler *handler.AuthHandler,
	transactionHandler *handler.TransactionHandler,
) []*Route {
	return []*Route{
		{
			Method:  echo.POST,
			Path:    "/login",
			Handler: authHandler.Login,
		},
		{
			Method:  echo.POST,
			Path:    "/transactions/webhook",
			Handler: transactionHandler.WebHookTransaction,
		},
	}
}

func PrivateRoutes(
	userHandler *handler.UserHandler,
	transactionHandler *handler.TransactionHandler,
) []*Route {
	return []*Route{
		{
			Method:  echo.GET,
			Path:    "/users",
			Handler: userHandler.GetAllUsers,
			Roles:   allRoles,
		},
		{
			Method:  echo.GET,
			Path:    "/users/:id",
			Handler: userHandler.GetUserByID,
			Roles:   allRoles,
		},
		{
			Method:  echo.POST,
			Path:    "/users",
			Handler: userHandler.CreateUser,
			Roles:   onlyEditor,
		},
		{
			Method:  echo.PUT,
			Path:    "/users/:id",
			Handler: userHandler.UpdateUser,
			Roles:   onlyEditor,
		},
		{
			Method:  echo.DELETE,
			Path:    "/users/:id",
			Handler: userHandler.DeleteUser,
			Roles:   onlyAdmin,
		},
		{
			Method:  echo.POST,
			Path:    "/transactions",
			Handler: transactionHandler.CreateOrder,
			Roles:   allRoles,
		},
	}
}
