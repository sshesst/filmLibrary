package test

//
//import (
//	"errors"
//	"filmLibrary/internal/models"
//	"filmLibrary/internal/service"
//	"filmLibrary/internal/storage"
//	"testing"
//
//	"github.com/DATA-DOG/go-sqlmock"
//)
//
//func TestAddMovieToDB(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	mockDB := storage.NewMockDB()
//
//	movie := models.Movie{
//		Title:       "Test Movie",
//		Description: "Test Description",
//		ReleaseDate: "2024-03-18",
//		Rating:      5,
//		Actors:      []models.Actor{{ID: 1}, {ID: 2}},
//	}
//
//	mock.ExpectQuery("SELECT COUNT(.+)").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
//	mock.ExpectExec("INSERT INTO movies").WillReturnResult(sqlmock.NewResult(1, 1))
//	mock.ExpectQuery("SELECT id FROM movies").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
//	mock.ExpectExec("INSERT INTO movie_actors").WillReturnResult(sqlmock.NewResult(1, 1))
//
//	err = service.AddMovieToDB(movie, mockDB)
//	if err != nil {
//		t.Errorf("unexpected error: %s", err)
//	}
//}
//
//func TestAddMovieToDB_Duplicate(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	mockDB := storage.NewMockDB()
//
//	movie := models.Movie{
//		Title:       "Test Movie",
//		Description: "Test Description",
//		ReleaseDate: "2024-03-18",
//		Rating:      5,
//		Actors:      []models.Actor{{ID: 1}, {ID: 2}},
//	}
//
//	mock.ExpectQuery("SELECT COUNT(.+)").WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
//
//	err = service.AddMovieToDB(movie, mockDB)
//	if err == nil || err.Error() != "такой фильм уже существует" {
//		t.Errorf("expected error: такой фильм уже существует, got: %s", err)
//	}
//}
//
//func TestAddMovieToDB_DBError(t *testing.T) {
//	db, mock, err := sqlmock.New()
//	if err != nil {
//		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
//	}
//	defer db.Close()
//
//	mockDB := storage.NewMockDB()
//
//	movie := models.Movie{
//		Title:       "Test Movie",
//		Description: "Test Description",
//		ReleaseDate: "2024-03-18",
//		Rating:      5,
//		Actors:      []models.Actor{{ID: 1}, {ID: 2}},
//	}
//
//	mock.ExpectQuery("SELECT COUNT(.+)").WillReturnError(errors.New("database error"))
//
//	err = service.AddMovieToDB(movie, mockDB)
//	if err == nil || err.Error() != "database error" {
//		t.Errorf("expected error: database error, got: %s", err)
//	}
//}
