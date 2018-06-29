package main

import (
	"go-test/src/infrastructure"
	"go-test/src/interfaces/webservices"
	"net/http"
	"time"

	"github.com/labstack/echo"
)

func setupRoutes(e *echo.Echo, c *config, auth *webservices.Auth, users *webservices.Users) {
	r := e.Group("", infrastructure.JSONHeadersMiddleware)
	r.Any("/", defaultHandler)
	r.POST("/login", auth.Login)

	user := r.Group("/users", infrastructure.JWTMiddleware([]byte(c.SigningKey)))
	{
		user.GET("", users.Get)
		user.POST("", users.Create)
		user.DELETE("/:id", users.Delete)
	}
}

func defaultHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":      http.StatusText(http.StatusOK),
		"server_time": time.Now(),
	})
}
