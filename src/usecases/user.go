package usecases

import (
	"errors"
	"go-test/src/domain"
)

// Predefined errors
var (
	ErrMissedRole       = errors.New("Role is required")
	ErrNotAvailableRole = errors.New("Not available role")
	ErrDeleteForbidden  = errors.New("Not enough permissions to delete users")
	ErrEmailIsTaken     = errors.New("Email is already in use")
)

// NewUserInteractor is a factory function,
// returns a new instance of the UserInteractor structure
func NewUserInteractor(repo domain.UserRepository) *UserInteractor {
	return &UserInteractor{repo}
}

// UserInteractor structure
type UserInteractor struct {
	repo domain.UserRepository
}

// User structure
type User struct {
	domain.User
}

// GetAll returns array of domain.User structures or error
func (i *UserInteractor) GetAll(scope string) (users []User, err error) {
	var domainUsers []domain.User
	if scope == domain.RoleAdmin {
		domainUsers, err = i.repo.GetAll()
		if err != nil {
			return
		}
	} else {
		domainUsers, err = i.repo.GetByRole(domain.RoleUser)
		if err != nil {
			return
		}
	}

	users = make([]User, len(domainUsers))
	for i, user := range domainUsers {
		users[i] = User{user}
	}

	return users, nil
}

// Create a new user
func (i *UserInteractor) Create(scope string, data map[string]interface{}) error {
	role, ok := data["role"]
	if !ok {
		return ErrMissedRole
	}
	if scope == domain.RoleUser && role != domain.RoleUser {
		return ErrNotAvailableRole
	}

	email, ok := data["email"]
	if ok {
		if user, err := i.repo.GetByEmail(email.(string)); err == nil && user.ID > 0 {
			return ErrEmailIsTaken
		}
	}

	user, err := domain.NewUser(data)
	if err != nil {
		return err
	}
	if _, err := i.repo.Store(user); err != nil {
		return err
	}

	return nil
}

// Delete a user
func (i *UserInteractor) Delete(scope string, id int) error {
	if scope != domain.RoleAdmin {
		return ErrDeleteForbidden
	}
	if err := i.repo.Delete(id); err != nil {
		return err
	}
	return nil
}
