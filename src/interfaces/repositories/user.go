package repositories

import (
	"go-test/src/domain"
)

// NewUserRepository is a factory function,
// returns an instance of UserRepository structure
func NewUserRepository(db DbHandler) *UserRepository {
	return &UserRepository{db}
}

// UserRepository structure
type UserRepository struct {
	db DbHandler
}

// GetByID retrieves user from storage by id
func (r *UserRepository) GetByID(id int) (domain.User, error) {
	q := `SELECT id, name, email, password, role FROM users WHERE id = ? LIMIT 1`
	rows, err := r.db.Query(q, id)
	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
		if err != nil {
			return domain.User{}, err
		}
	}

	return user, nil
}

// GetByEmail retrieves user from storage by email
func (r *UserRepository) GetByEmail(email string) (domain.User, error) {
	q := `SELECT id, name, email, password, role FROM users WHERE email = ? LIMIT 1`
	rows, err := r.db.Query(q, email)
	if err != nil {
		return domain.User{}, err
	}

	user := domain.User{}
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
		if err != nil {
			return domain.User{}, err
		}
	}

	return user, nil
}

// GetAll retrieves all users from storage
func (r *UserRepository) GetAll() ([]domain.User, error) {
	q := `SELECT id, name, email, password, role FROM users`
	rows, err := r.db.Query(q)
	if err != nil {
		return []domain.User{}, err
	}

	users := []domain.User{}
	for rows.Next() {
		user := domain.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
		if err != nil {
			return []domain.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetByRole retrieves list of users from storage by role
func (r *UserRepository) GetByRole(role string) ([]domain.User, error) {
	q := `SELECT id, name, email, password, role FROM users WHERE role = ?`
	rows, err := r.db.Query(q, role)
	if err != nil {
		return []domain.User{}, err
	}

	users := []domain.User{}
	for rows.Next() {
		user := domain.User{}
		err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Role)
		if err != nil {
			return []domain.User{}, err
		}
		users = append(users, user)
	}

	return users, nil
}

// Store user to storage
func (r *UserRepository) Store(user domain.User) (id int64, err error) {
	q := `INSERT INTO users (name, email, password, role) VALUES (?, ?, ?, ?)`
	res, err := r.db.Execute(q, user.Name, user.Email, user.Password, user.Role)
	if err != nil {
		return
	}
	id, err = res.LastInsertId()
	if err != nil {
		return
	}
	return id, nil
}

// Delete user from storage
func (r *UserRepository) Delete(id int) error {
	q := `DELETE FROM users WHERE id = ?`
	_, err := r.db.Execute(q, id)
	if err != nil {
		return err
	}
	return nil
}
