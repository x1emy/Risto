package services

import (
	"booking-service/internal/models"
	"booking-service/internal/repositories"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Ключ для подписи JWT токенов
var jwtKey = []byte("your_secret_key")

type AuthService struct {
	userRepository repositories.UserRepository // Используем интерфейс
}

func NewAuthService(userRepo repositories.UserRepository) *AuthService {
	return &AuthService{userRepository: userRepo}
}

// Claims структура для токена
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Генерация токена
func (s *AuthService) GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "booking-service",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Создание пользователя
func (s *AuthService) CreateUser(username, password string) (*models.User, error) {
	// Логирование входящих данных
	log.Println("Попытка создания пользователя:", username)

	// Хэшируем пароль
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Ошибка хэширования пароля:", err)
		return nil, fmt.Errorf("ошибка хэширования пароля: %v", err)
	}

	// Создаем пользователя
	user := &models.User{
		Username: username,
		Password: string(hashedPassword),
	}

	// Сохраняем пользователя в репозитории
	err = s.userRepository.CreateUser(user)
	if err != nil {
		log.Println("Ошибка при сохранении пользователя:", err)
		return nil, err
	}

	// Логируем успешное создание пользователя
	log.Println("Пользователь успешно создан:", user)

	return user, nil
}

// Аутентификация
func (s *AuthService) Authenticate(username, password string) (string, error) {
	user, err := s.userRepository.GetUserByUsername(username)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", fmt.Errorf("неверный пароль")
	}

	token, err := s.GenerateJWT(username)
	if err != nil {
		return "", err
	}

	return token, nil
}

// Проверка токена
func CheckAuth(r *http.Request) (string, error) {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return "", fmt.Errorf("отсутствует токен")
	}

	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("неожиданный метод подписи")
		}
		return jwtKey, nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("недействительный токен")
	}

	return claims.Username, nil
}
