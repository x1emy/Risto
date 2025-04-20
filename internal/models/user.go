package models

// User - структура для пользователя
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"` // Пароль не будет возвращаться в ответе
}
