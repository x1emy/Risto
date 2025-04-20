package repositories

import (
	"database/sql"
	"log"
)

// Table - структура для столика
type Table struct {
	ID          int    `json:"id"`
	TableNumber int    `json:"table_number"`
	Seats       int    `json:"seats"`
	Status      string `json:"status"`
}

// BookingRepository - репозиторий для работы с бронированиями
type BookingRepository struct {
	db *sql.DB
}

func connectToDB() (*sql.DB, error) {
	connStr := "postgres://postgres:1202@localhost:5432/booking_db?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	// Проверка подключения
	if err := db.Ping(); err != nil {
		return nil, err
	}
	log.Println("Подключение к базе данных успешно!")
	return db, nil
}

// NewBookingRepository создает новый репозиторий
func NewBookingRepository(db *sql.DB) *BookingRepository {
	return &BookingRepository{
		db: db,
	}
}

// GetAllTables получает все доступные столики
func (r *BookingRepository) GetAllTables() ([]Table, error) {
	rows, err := r.db.Query("SELECT id, table_number, seats, status FROM tables WHERE status = 'free'")
	if err != nil {
		log.Println("Ошибка при выполнении запроса:", err)
		return nil, err
	}
	defer rows.Close()

	var tables []Table
	for rows.Next() {
		var table Table
		if err := rows.Scan(&table.ID, &table.TableNumber, &table.Seats, &table.Status); err != nil {
			log.Println("Ошибка при сканировании строки:", err)
			return nil, err
		}
		tables = append(tables, table)
	}

	if len(tables) == 0 {
		log.Println("Нет доступных столиков")
	}

	return tables, nil
}

// BookTable забронировать столик
func (r *BookingRepository) BookTable(tableID int, guestName string, bookingTime string) error {
	_, err := r.db.Exec("INSERT INTO bookings (table_id, guest_name, booking_time) VALUES ($1, $2, $3)", tableID, guestName, bookingTime)
	return err
}
