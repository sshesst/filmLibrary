package controllers

import (
	"context"
	"encoding/json"
	database "filmLibrary"
	"filmLibrary/internal/models"
	"net/http"
	"strings"
	"time"
)

// SearchMoviesByActor выполняет поиск фильмов по фрагменту имени актёра
func SearchMoviesByActor(w http.ResponseWriter, r *http.Request) {
	// Получаем параметр запроса для поиска
	query := r.URL.Query().Get("query")

	// Если запрос пустой, возвращаем пустой результат
	if query == "" {
		http.Error(w, "Query parameter is required", http.StatusBadRequest)
		return
	}

	// Выполняем запрос к базе данных для поиска фильмов по имени актёра
	movies, err := searchMoviesByActorFromDB(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем результаты поиска в виде JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// searchMoviesByActorFromDB выполняет поиск фильмов в базе данных по фрагменту имени актёра
func searchMoviesByActorFromDB(query string) ([]models.Movie, error) {
	pool, err := database.GetPool()
	if err != nil {
		return nil, err
	}
	defer pool.Close()

	// Выполняем запрос к базе данных для поиска фильмов по фрагменту имени актёра
	rows, err := pool.Query(context.Background(), `
		SELECT m.id, m.title, m.description, m.release_date, m.rating
		FROM movies m
		JOIN movie_actors ma ON m.id = ma.movie_id
		JOIN actors a ON ma.actor_id = a.id
		WHERE LOWER(a.name) LIKE LOWER($1)
	`, "%"+strings.ToLower(query)+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		var releaseDate time.Time
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &releaseDate, &movie.Rating)
		if err != nil {
			return nil, err
		}
		movie.ReleaseDate = releaseDate.Format("2006-01-02")
		movies = append(movies, movie)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return movies, nil
}
