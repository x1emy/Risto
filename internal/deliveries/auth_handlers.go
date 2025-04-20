// deliveries/auth_handlers.go
package deliveries

import (
	"booking-service/internal/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RegisterAuthHandlers(r *mux.Router, authService *services.AuthService) {
	// 🔁 Добавим поддержку OPTIONS (Preflight) — обязательно для CORS
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}).Methods("OPTIONS")

	// ✅ Обработка POST-запроса регистрации
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		var userInput struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}

		if err := json.NewDecoder(r.Body).Decode(&userInput); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		log.Println("Полученные данные для регистрации:", userInput)

		user, err := authService.CreateUser(userInput.Username, userInput.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}).Methods("POST")
}
