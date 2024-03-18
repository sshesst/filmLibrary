package controllers

import (
	"context"
	"encoding/json"
	"filmLibrary/internal/models"
	database "filmLibrary/internal/storage"
	"filmLibrary/pkg/logging"
	"net/http"
	"strings"
	"time"
)

func SearchMoviesByName(w http.ResponseWriter, r *http.Request, logger logging.Logger) {
	query := r.URL.Query().Get("query")

	if query == "" {
		http.Error(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	movies, err := searchMoviesFromDB(query, logger)
	if err != nil {
		logger.Error("Ошибка поиска фильмов в БД:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func searchMoviesFromDB(query string, logger logging.Logger) ([]models.Movie, error) {
	pool, err := database.GetPool()
	if err != nil {
		logger.Error("Ошибка получения пула соединений:", err)
		return nil, err
	}
	defer pool.Close()

	rows, err := pool.Query(context.Background(), `
		SELECT id, title, description, release_date, rating
		FROM movies
		WHERE LOWER(title) LIKE LOWER($1)
	`, "%"+strings.ToLower(query)+"%")
	if err != nil {
		logger.Error("Ошибка выполнения SQL-запроса:", err)
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		var releaseDate time.Time
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &releaseDate, &movie.Rating)
		if err != nil {
			logger.Error("Ошибка сканирования результата:", err)
			return nil, err
		}
		movie.ReleaseDate = releaseDate.Format("2006-01-02")
		movies = append(movies, movie)
	}
	if err := rows.Err(); err != nil {
		logger.Error("Ошибка перебора результатов:", err)
		return nil, err
	}

	return movies, nil
}
