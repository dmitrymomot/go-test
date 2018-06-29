package webservices

import (
	"go-test/src/usecases"
	"net/http"

	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
)

// JWTHandler interface
type JWTHandler interface {
	Make(userID int, scope string) (token string, err error)
}

// NewAuthWebservice is a factory function,
// returns an instance of Auth webservice structure
func NewAuthWebservice(i *usecases.AuthInteractor, l Logger, h JWTHandler) *Auth {
	return &Auth{i, l, h}
}

// Auth web service structure
type Auth struct {
	interactor *usecases.AuthInteractor
	logger     Logger
	jwtHandler JWTHandler
}

// Login is a route handler,
// returns auth token
func (ws *Auth) Login(c echo.Context) error {
	form := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{}
	v := govalidator.New(govalidator.Options{
		Request: c.Request(),
		Data:    &form,
		Rules: govalidator.MapData{
			"email":    []string{"required", "email"},
			"password": []string{"required", "min:5", "max:50"},
		},
	})
	if err := v.ValidateJSON(); len(err) > 0 {
		if _, ok := err["_error"]; ok {
			return ErrMalformedJSON
		}
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user, err := ws.interactor.Login(form.Email, form.Password)
	if err != nil {
		switch err {
		case usecases.ErrInvalidCredentials:
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	token, err := ws.jwtHandler.Make(user.ID, user.Role)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, response{Data: data{"token": token}})
}
