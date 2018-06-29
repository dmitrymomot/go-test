package domain

import (
	"errors"
	"log"

	"github.com/jmcvetta/randutil"
	"github.com/mitchellh/mapstructure"
	"golang.org/x/crypto/bcrypt"
)

// Predefined roles
const (
	RoleUser  = "user"
	RoleAdmin = "admin"
)

// Predefined errors
var (
	ErrUndefinedRole      = errors.New("Undefined user role")
	ErrCouldNotCreateUser = errors.New("Could not create user with provided data")
	ErrMissedEmail        = errors.New("Email is required")
)

// UserRepository interface
type UserRepository interface {
	GetByID(id int) (User, error)
	GetByEmail(email string) (User, error)
	GetAll() ([]User, error)
	GetByRole(role string) ([]User, error)
	Store(user User) (id int64, err error)
	Delete(id int) error
}

// NewUser is a factory function,
// returns a new instance or User structure
func NewUser(data map[string]interface{}) (User, error) {
	user := User{}
	mapToStruct(data, &user)

	if user.Role != RoleUser && user.Role != RoleAdmin {
		return User{}, ErrUndefinedRole
	}

	if user.Email == "" {
		return User{}, ErrMissedEmail
	}

	if password, ok := data["password"]; ok {
		user.Password = password.(string)
	} else {
		user.Password = randString(10)
	}
	user.Password = hashString(user.Password)

	return user, nil
}

// User structure
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email"`
	Password string `json:"-"`
	Role     string `json:"role"`
}

// Hash hashes string
func hash(s string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// HashString hashes string
func hashString(s string) string {
	hash, err := hash(s)
	if err != nil {
		log.Println(err)
		return ""
	}
	return hash
}

func randString(n int) string {
	str, _ := randutil.AlphaString(n)
	return str
}

func mapToStruct(data map[string]interface{}, output interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:  "json",
		Metadata: nil,
		Result:   output,
	})
	if err != nil {
		return err
	}
	if err := decoder.Decode(data); err != nil {
		return err
	}
	return nil
}
