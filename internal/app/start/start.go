package start

import (
	"booking-service/internal/deliveries"
	"booking-service/internal/repositories"
	"booking-service/internal/services"
	"database/sql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"path/filepath"
)

// Функция для обработки сервера
func RunServer(db *sql.DB) {
	r := mux.NewRouter()

	// Репозитории
	authRepo := repositories.NewAuthRepository(db)
	bookingRepo := repositories.NewBookingRepository(db)

	// Сервисы
	authService := services.NewAuthService(authRepo)
	bookingService := services.NewBookingService(bookingRepo)

	// Разрешаем CORS
	r.Use(corsMiddleware)

	// Обработчики
	deliveries.RegisterAuthHandlers(r, authService)
	deliveries.RegisterBookingHandlers(r, bookingService) // Для получения данных о столиках

	// Абсолютный путь к папке dist/booking-frontend, где собран Angular
	publicDir, err := filepath.Abs("booking-frontend/dist/booking-frontend")
	if err != nil {
		log.Fatalf("Не удалось определить путь к папке фронтенда: %v", err)
	}

	// Главная страница — отправка собранного HTML из Angular
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Запрос на главную страницу")
		_, err := bookingService.GetAvailableTables()
		if err != nil {
			http.Error(w, "Ошибка получения данных о столиках", http.StatusInternalServerError)
			return
		}
		// Отправляем собранный HTML файл
		http.ServeFile(w, r, filepath.Join(publicDir, "index.html"))
	})

	// Статические файлы (JS, CSS и т.д.)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(filepath.Join(publicDir, "static")))))

	// Перенаправление всех прочих маршрутов на index.html для поддержки маршрутизации Angular
	r.HandleFunc("/{rest:.*}", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(publicDir, "index.html"))
	})

	log.Println("Сервер запущен на :8080")
	http.ListenAndServe(":8080", r)
}

// CORS middleware для разрешения запросов с фронтенда
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Разрешаем доступ с порта 4200 (фронтенд Angular)
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200") // Указываем конкретный источник (origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Обработка preflight-запросов (OPTIONS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Продолжаем обработку запроса
		next.ServeHTTP(w, r)
	})
}
