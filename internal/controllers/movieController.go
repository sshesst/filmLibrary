package controllers

import (
	"encoding/json"
	"filmLibrary/internal/handler"
	"filmLibrary/internal/models"
	"filmLibrary/internal/service"
	"filmLibrary/pkg/logging"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func AddMovie(w http.ResponseWriter, r *http.Request, logger logging.Logger) {
	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		logger.Error("Ошибка декодирования JSON:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handler.MovieHandler(movie, logger)
	if err != nil {
		logger.Error("Ошибка проверки лимита введённых символов в запросе", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.AddMovieToDB(movie, logger)
	if err != nil {
		logger.Error("Ошибка добавления фильма в БД:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("Фильм успешно добавлен")
	fmt.Fprintf(w, "Фильм успешно добавлен")
}

func UpdateMovie(w http.ResponseWriter, r *http.Request, logger logging.Logger) {
	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		logger.Error("Ошибка декодирования JSON:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.UpdateMovieInDB(movie, logger)
	if err != nil {
		logger.Error("Ошибка обновления фильма в БД:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("Информация о фильме успешно обновлена")
	fmt.Fprintf(w, "Информация о фильме успешно обновлена")
}

func DeleteMovie(w http.ResponseWriter, r *http.Request, logger logging.Logger) {
	vars := mux.Vars(r)
	id, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		logger.Error("Ошибка декодирования JSON:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.DeleteMovieFromDB(uint(id), logger)
	if err != nil {
		logger.Error("Ошибка удаления фильма из БД:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	logger.Info("Информация о фильме успешно удалена")
	fmt.Fprintf(w, "Информация о фильме успешно удалена")
}
