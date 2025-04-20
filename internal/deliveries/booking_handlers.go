package deliveries

import (
	"booking-service/internal/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func RegisterBookingHandlers(r *mux.Router, bookingService *services.BookingService) {
	// 🔓 Публичный маршрут — доступные столики без авторизации
	r.HandleFunc("/api/tables", func(w http.ResponseWriter, r *http.Request) {
		tables, err := bookingService.GetAvailableTables()
		if err != nil {
			log.Println("Ошибка при получении столиков:", err)
			http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(tables); err != nil {
			log.Println("Ошибка при кодировании JSON:", err)
			http.Error(w, "Ошибка кодирования ответа", http.StatusInternalServerError)
		}
	}).Methods("GET")

	// 🔐 Получение столиков только для авторизованных
	r.HandleFunc("/tables", func(w http.ResponseWriter, r *http.Request) {
		_, err := services.CheckAuth(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tables, err := bookingService.GetAvailableTables()
		if err != nil {
			log.Println("Ошибка при получении доступных столиков:", err)
			http.Error(w, "Ошибка получения данных", http.StatusInternalServerError)
			return
		}

		if len(tables) == 0 {
			log.Println("Нет доступных столиков")
			http.Error(w, "Нет доступных столиков", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(tables); err != nil {
			log.Println("Ошибка при кодировании ответа:", err)
			http.Error(w, "Ошибка кодирования ответа", http.StatusInternalServerError)
		}
	}).Methods("GET")

	// Приветствие
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Добро пожаловать в систему бронирования столиков!"))
	}).Methods("GET")

	// POST /booking — Бронирование
	r.HandleFunc("/booking", func(w http.ResponseWriter, r *http.Request) {
		var booking struct {
			TableID     int    `json:"table_id"`
			GuestName   string `json:"guest_name"`
			BookingTime string `json:"booking_time"`
		}

		if err := json.NewDecoder(r.Body).Decode(&booking); err != nil {
			http.Error(w, "Неверный формат данных", http.StatusBadRequest)
			return
		}

		err := bookingService.BookTable(booking.TableID, booking.GuestName, booking.BookingTime)
		if err != nil {
			http.Error(w, "Ошибка бронирования", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}).Methods("POST")
}
