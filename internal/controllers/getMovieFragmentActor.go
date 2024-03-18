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

func SearchMoviesByActor(w http.ResponseWriter, r *http.Request, logger logging.Logger) {
	query := r.URL.Query().Get("query")

	if query == "" {
		logger.Error("Query parameter is required")
		http.Error(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	movies, err := searchMoviesByActorFromDB(query, logger)
	if err != nil {
		logger.Error("Ошибка поиска фильмов по актеру в БД:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func searchMoviesByActorFromDB(query string, logger logging.Logger) ([]models.Movie, error) {
	pool, err := database.GetPool()
	if err != nil {
		logger.Error("Ошибка получения пула соединений:", err)
		return nil, err
	}
	defer pool.Close()

	rows, err := pool.Query(context.Background(), `
		SELECT m.id, m.title, m.description, m.release_date, m.rating
		FROM movies m
		JOIN movie_actors ma ON m.id = ma.movie_id
		JOIN actors a ON ma.actor_id = a.id
		WHERE LOWER(a.name) LIKE LOWER($1)
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
