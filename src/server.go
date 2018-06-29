package main

import (
	"go-test/src/interfaces/webservices"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// NewServer is a factory function, returns a new instance of Server structure
func NewServer(config *config, auth *webservices.Auth, users *webservices.Users) *Server {
	e := echo.New()
	e.Debug = config.Debug

	// Set default error handler
	e.HTTPErrorHandler = defaultHTTPErrorHandler

	// Will run before middlewares
	e.Pre(middleware.RemoveTrailingSlash())

	// Middleware
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "${time_rfc3339} [${status}] ${method}:${uri}, Latency: ${latency_human}\n",
	}))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.SecureWithConfig(middleware.SecureConfig{
		Skipper:               middleware.DefaultSkipper,
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		XFrameOptions:         "DENY",
		ContentSecurityPolicy: "default-src 'self'",
	}))

	return &Server{e, config, auth, users}
}

// Server struct
type Server struct {
	e       *echo.Echo
	c       *config
	authWs  *webservices.Auth
	usersWs *webservices.Users
}

// Run Server
func (s *Server) Run() {
	setupRoutes(s.e, s.c, s.authWs, s.usersWs)
	s.e.Use(middleware.NonWWWRedirect())
	s.e.Logger.Fatal(s.e.Start(s.c.ListenAddress))
}

func defaultHTTPErrorHandler(err error, c echo.Context) {
	e := struct {
		Code   int         `json:"code"`
		Title  string      `json:"title"`
		Detail interface{} `json:"detail,omitempty"`
	}{}
	if he, ok := err.(*echo.HTTPError); ok {
		e.Code = he.Code
		e.Title = http.StatusText(he.Code)
		e.Detail = he.Message
		if he.Internal != nil {
			c.Logger().Error(he.Internal)
		}
	} else {
		e.Code = http.StatusInternalServerError
		e.Title = http.StatusText(e.Code)
		e.Detail = err
	}
	checkErr(c.JSON(e.Code, map[string]interface{}{"error": e}))
}
