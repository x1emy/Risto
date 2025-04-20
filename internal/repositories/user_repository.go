package repositories

import (
	"booking-service/internal/models"
	"database/sql"
	"errors"
)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByUsername(username string) (*models.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) error {
	_, err := r.db.Exec(
		"INSERT INTO users (username, password) VALUES ($1, $2)",
		user.Username,
		user.Password,
	)
	return err
}

func (r *userRepository) GetUserByUsername(username string) (*models.User, error) {
	row := r.db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
