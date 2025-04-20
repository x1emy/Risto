package config

import (√
	"log"
	"os"
)

var (
	DBHost     = os.Getenv("DB_HOST")
	DBPort     = os.Getenv("DB_PORT")
	DBUser     = os.Getenv("DB_USER")
	DBPassword = os.Getenv("DB_PASSWORD")
	DBName     = os.Getenv("DB_NAME")
)

func LoadConfig() {
	log.Println("DB_HOST:", DBHost)
	log.Println("DB_PORT:", DBPort)
	log.Println("DB_USER:", DBUser)
	log.Println("DB_PASSWORD:", DBPassword)
	log.Println("DB_NAME:", DBName)

	if DBHost == "" || DBPort == "" || DBUser == "" || DBPassword == "" || DBName == "" {
		log.Fatal("❌ Недостающие переменные окружения для подключения к БД")
	}
}
