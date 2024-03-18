package storage

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

// DB представляет интерфейс базы данных
type DB interface {
	GetPool() (*pgxpool.Pool, error)
	CreateTables(pool *pgxpool.Pool) error
}

// MockDB представляет мокированную реализацию интерфейса базы данных для тестов
type MockDB struct{}

// GetPool возвращает мокированный пул соединений
func (m *MockDB) GetPool() (*pgxpool.Pool, error) {
	// Возвращаем мокированный пул соединений для тестов
	return nil, nil // Замените на вашу реализацию
}

// CreateTables создает мокированные таблицы
func (m *MockDB) CreateTables(pool *pgxpool.Pool) error {
	// Здесь вы можете имитировать создание таблиц в мокированной базе данных
	return nil // Замените на вашу реализацию
}

// NewMockDB возвращает новый экземпляр объекта базы данных для мокирования
func NewMockDB() DB {
	return &MockDB{}
}
