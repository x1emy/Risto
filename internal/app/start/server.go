package start

import (
	"booking-service/internal/models"
	"booking-service/internal/services"
	"encoding/json"
	"net/http"
)

// Функция обработки регистрации пользователя
func RegisterUser(w http.ResponseWriter, r *http.Request, authService *services.AuthService) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Ошибка чтения данных", http.StatusBadRequest)
		return
	}

	// Создание пользователя
	createdUser, err := authService.CreateUser(user.Username, user.Password)
	if err != nil {
		http.Error(w, "Ошибка создания пользователя", http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ с данными нового пользователя
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(createdUser)
}
