package repository

import (
	"database/sql"

	"cambria-bank/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (r UserRepository) CreateUser(user models.User) (int, error) {

	query := `
		INSERT INTO users (name, email, password, phone, role)
		VALUES (?, ?, ?, ?, ?)
	`

	result, err := r.DB.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Phone,
		user.Role,
	)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r UserRepository) GetUserByEmail(email string) (models.User, error) {

	var user models.User

	query := `
		SELECT id, name, email, password, phone, role
		FROM users
		WHERE email = ?
	`

	err := r.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Phone,
		&user.Role,
	)

	return user, err
}
