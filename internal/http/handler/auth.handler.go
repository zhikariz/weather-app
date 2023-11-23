package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/zhikariz/weather-app/internal/http/validator"
	"github.com/zhikariz/weather-app/internal/service"
)

type AuthHandler struct {
	loginService service.LoginUseCase
	tokenService service.TokenUseCase
}

func NewAuthHandler(
	loginService service.LoginUseCase,
	tokenService service.TokenUseCase,
) *AuthHandler {
	return &AuthHandler{
		loginService: loginService,
		tokenService: tokenService,
	}
}

func (h *AuthHandler) Login(ctx echo.Context) error {
	var input struct {
		Email    string `json:"email" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	if err := ctx.Bind(&input); err != nil {
		return ctx.JSON(http.StatusBadRequest, validator.ValidatorErrors(err))
	}

	user, err := h.loginService.Login(ctx.Request().Context(), input.Email, input.Password)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	accessToken, err := h.tokenService.GenerateAccessToken(ctx.Request().Context(), user)

	if err != nil {
		return ctx.JSON(http.StatusUnprocessableEntity, err)
	}

	data := map[string]string{
		"access_token": accessToken,
	}

	return ctx.JSON(http.StatusOK, data)
}
