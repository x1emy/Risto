package main

import (
	"booking-service/internal/app/config"
	"booking-service/internal/app/connections"
	"booking-service/internal/app/start"
	"log"
)

func main() {
	// Загружаем конфиг из переменных окружения
	config.LoadConfig()

	// Подключаемся к базе данных
	db, err := connections.InitDB()
	if err != nil {
		log.Fatal("Ошибка подключения к БД: ", err)
	}

	// Запуск HTTP сервера
	start.RunServer(db)
}
