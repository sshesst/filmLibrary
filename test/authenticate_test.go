package test

//import (
//	"filmLibrary/internal/models"
//	"filmLibrary/internal/service"
//	"testing"
//
//	"github.com/DATA-DOG/go-sqlmock"
//)
//
//// TestAuthenticate проверяет функцию Authenticate.
//func TestAuthenticate(t *testing.T) {
//	// Создание мока для SQL-соединения
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("Ошибка создания мока: %v", err)
//	}
//	defer db.Close()
//
//	// Настройка ожидаемого запроса и результата мока
//	rows := sqlmock.NewRows([]string{"id", "is_admin"}).AddRow(1, true)
//	mock.ExpectQuery("SELECT id, is_admin FROM users WHERE username = ? AND password = ?").
//		WithArgs("testuser", "testpassword").
//		WillReturnRows(rows)
//
//	// Подмена настоящего соединения моком
//	service.SetDB(db)
//
//	// Выполнение функции
//	user, err := service.Authenticate("testuser", "testpassword")
//
//	// Проверка результата
//	if err != nil {
//		t.Errorf("Ошибка: %v", err)
//	}
//	expectedUser := models.User{ID: 1, IsAdmin: true}
//	if user != expectedUser {
//		t.Errorf("Неправильный пользователь. Получено: %v, ожидается: %v", user, expectedUser)
//	}
//
//	// Проверка вызовов мока
//	if err := mock.ExpectationsWereMet(); err != nil {
//		t.Errorf("Ожидаемые вызовы не выполнены: %s", err)
//	}
//}
