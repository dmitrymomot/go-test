package webservices

import (
	"go-test/src/usecases"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/thedevsaddam/govalidator"
)

// NewUsersWebservice is a factory function,
// returns an instance of Users webservice structure
func NewUsersWebservice(i *usecases.UserInteractor, l Logger) *Users {
	return &Users{i, l}
}

// Users web service structure
type Users struct {
	interactor *usecases.UserInteractor
	logger     Logger
}

// Get is a route handler,
// returns list of users
func (ws *Users) Get(c echo.Context) error {
	scope, err := getScopeFromContext(c)
	if err != nil {
		return err
	}
	users, err := ws.interactor.GetAll(scope)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"data": users})
}

// Create is a route handler,
// stores a new user
func (ws *Users) Create(c echo.Context) error {
	scope, err := getScopeFromContext(c)
	if err != nil {
		return err
	}
	data := make(map[string]interface{}, 0)
	v := govalidator.New(govalidator.Options{
		Request: c.Request(),
		Data:    &data,
		Rules: govalidator.MapData{
			"name":     []string{"required", "min:2", "max:250"},
			"email":    []string{"required", "email"},
			"password": []string{"required", "min:5", "max:50"},
			"role":     []string{"required"},
		},
	})
	if err := v.ValidateJSON(); len(err) > 0 {
		if _, ok := err["_error"]; ok {
			return ErrMalformedJSON
		}
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := ws.interactor.Create(scope, data); err != nil {
		switch err {
		case usecases.ErrMissedRole, usecases.ErrNotAvailableRole, usecases.ErrEmailIsTaken:
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{"data": true})
}

// Delete is a route handler,
// deletes a user from storage
func (ws *Users) Delete(c echo.Context) error {
	scope, err := getScopeFromContext(c)
	if err != nil {
		return err
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return ErrWrongIDParam
	}
	if err := ws.interactor.Delete(scope, id); err != nil {
		switch err {
		case usecases.ErrDeleteForbidden:
			return echo.NewHTTPError(http.StatusForbidden, err.Error())
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"data": true})
}

func getScopeFromContext(c echo.Context) (string, error) {
	user := c.Get("JWT").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	scope, ok := claims["scope"].(string)
	if !ok {
		return "", ErrForbidden
	}
	return scope, nil
}
