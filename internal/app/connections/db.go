package connections

import (
	"booking-service/internal/app/config"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func InitDB() (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД: ", err)
		return nil, err
	}
	return db, nil
}
