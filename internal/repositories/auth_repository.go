package repositories

import (
	"booking-service/internal/models"
	"database/sql"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type AuthRepository struct {
	db *sql.DB
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

// CreateUser создает пользователя с хэшированным паролем (соответствует UserRepository интерфейсу)
func (r *AuthRepository) CreateUser(user *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("ошибка хэширования пароля: %v", err)
	}

	_, err = r.db.Exec(
		"INSERT INTO users (username, password) VALUES ($1, $2)", // Используем $1 и $2 для PostgreSQL
		user.Username,
		string(hashedPassword), // Хэшированный пароль
	)
	return err
}

// GetUserByUsername получает пользователя по имени
func (r *AuthRepository) GetUserByUsername(username string) (*models.User, error) {
	query := "SELECT id, username, password FROM users WHERE username = ?"
	row := r.db.QueryRow(query, username)

	var user models.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
