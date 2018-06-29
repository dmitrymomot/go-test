package usecases

import (
	"errors"
	"go-test/src/domain"

	"golang.org/x/crypto/bcrypt"
)

// Predefined errors
var (
	ErrInvalidCredentials = errors.New("Wrong email or password")
)

// NewAuthInteractor is a factory function,
// returns a new instance of the AuthInteractor structure
func NewAuthInteractor(repo domain.UserRepository) *AuthInteractor {
	return &AuthInteractor{repo}
}

// AuthInteractor structure
type AuthInteractor struct {
	repo domain.UserRepository
}

// Login usecase handler,
// returns instance of domain.User structure or error
func (i *AuthInteractor) Login(email, password string) (User, error) {
	user, err := i.repo.GetByEmail(email)
	if err != nil {
		return User{}, ErrInvalidCredentials
	}
	if !checkHash(user.Password, password) {
		return User{}, ErrInvalidCredentials
	}
	return User{user}, nil
}

// checkHash compares hash with string
func checkHash(hash, str string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}
