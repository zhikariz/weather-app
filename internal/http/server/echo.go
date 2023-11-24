package server

import (
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/zhikariz/weather-app/common"
	"github.com/zhikariz/weather-app/internal/config"
	"github.com/zhikariz/weather-app/internal/http/binder"
	"github.com/zhikariz/weather-app/internal/http/router"
)

type Server struct {
	*echo.Echo
}

func NewServer(
	cfg *config.Config,
	binder *binder.Binder,
	publicRoutes, privateRoutes []*router.Route) *Server {
	e := echo.New()
	e.HideBanner = true
	e.Binder = binder

	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.CORS(),
		session.Middleware(sessions.NewCookieStore([]byte(cfg.Session.SecretKey))),
	)

	v1 := e.Group("/api/v1")

	for _, public := range publicRoutes {
		v1.Add(public.Method, public.Path, public.Handler)
	}

	for _, private := range privateRoutes {
		// v1.Add(private.Method, private.Path, private.Handler, JWTProtected(cfg.JWT.SecretKey), SessionProtected())
		v1.Add(private.Method, private.Path, private.Handler, SessionProtected())
	}

	e.GET("/ping", func(c echo.Context) error {
		return c.String(200, "pong")
	})

	return &Server{e}
}

func JWTProtected(secretKey string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(common.JwtCustomClaims)
		},
		SigningKey: []byte(secretKey),
	})
}

func SessionProtected() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			sess, _ := session.Get("auth-sessions", ctx)
			if sess.Values["token"] == nil {
				return ctx.JSON(http.StatusUnauthorized, map[string]string{"error": "silahkan login terlebih dahulu"})
			}
			ctx.Set("user", sess.Values["token"])
			return next(ctx)
		}
	}
}
