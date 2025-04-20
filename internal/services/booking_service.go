package services

import (
	"booking-service/internal/repositories"
	"fmt"
)

// BookingService - сервис для бронирования столиков
type BookingService struct {
	repository *repositories.BookingRepository
}

// NewBookingService создает новый сервис бронирования
func NewBookingService(repo *repositories.BookingRepository) *BookingService {
	return &BookingService{
		repository: repo,
	}
}

// GetAvailableTables возвращает список доступных столиков
func (s *BookingService) GetAvailableTables() ([]repositories.Table, error) {
	tables, err := s.repository.GetAllTables()
	if err != nil {
		return nil, fmt.Errorf("не удалось получить доступные столики: %v", err)
	}
	return tables, nil
}

// BookTable бронирует столик для гостя
func (s *BookingService) BookTable(tableID int, guestName string, bookingTime string) error {
	// Вы можете добавить логику проверки доступности столика, прежде чем бронировать его
	err := s.repository.BookTable(tableID, guestName, bookingTime)
	if err != nil {
		return fmt.Errorf("не удалось забронировать столик: %v", err)
	}
	return nil
}
