package repositories

import "go-test/src/domain"

// Seed bootstrap data into storage
func Seed(ur domain.UserRepository) error {
	user, err := domain.NewUser(map[string]interface{}{
		"name":     "Test User",
		"email":    "user@testapp.loc",
		"password": "12345",
		"role":     domain.RoleUser,
	})
	if err != nil {
		return err
	}
	if _, err = ur.Store(user); err != nil {
		return err
	}

	admin, err := domain.NewUser(map[string]interface{}{
		"name":     "Test Admin",
		"email":    "admin@testapp.loc",
		"password": "12345",
		"role":     domain.RoleAdmin,
	})
	if err != nil {
		return err
	}
	if _, err = ur.Store(admin); err != nil {
		return err
	}

	return nil
}
