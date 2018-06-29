package webservices

import (
	"net/http"

	"github.com/labstack/echo"
)

// Predefined errors
var (
	ErrMalformedJSON = echo.NewHTTPError(http.StatusBadRequest, "Malformed JSON")
	ErrForbidden     = echo.NewHTTPError(http.StatusForbidden, "Access Denied")
	ErrWrongIDParam  = echo.NewHTTPError(http.StatusBadRequest, "ID parameter must be a number")
)
