package controllers

import (
	"context"
	"encoding/json"
	"filmLibrary/internal/models"
	database "filmLibrary/internal/storage"
	"filmLibrary/pkg/logging"
	"net/http"
	"time"
)

func GetMovies(w http.ResponseWriter, r *http.Request, logger logging.Logger) {
	sortBy := r.URL.Query().Get("sort_by")
	order := r.URL.Query().Get("order")

	if sortBy == "" {
		sortBy = "rating"
	}

	if order == "" {
		order = "desc"
	}

	sqlQuery := "SELECT m.id, m.title, m.description, m.release_date, m.rating, ma.actor_id, a.name, a.gender, a.birthdate " +
		"FROM movies m " +
		"JOIN movie_actors ma ON m.id = ma.movie_id " +
		"JOIN actors a ON ma.actor_id = a.id " +
		"ORDER BY " + sortBy + " " + order

	pool, err := database.GetPool()
	if err != nil {
		logger.Error("Ошибка получения пула соединений:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer pool.Close()

	rows, err := pool.Query(context.Background(), sqlQuery)
	if err != nil {
		logger.Error("Ошибка выполнения SQL-запроса:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
		var movie models.Movie
		var releaseDate time.Time
		var actorID uint
		var name string
		var gender string
		var birthdate time.Time
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &releaseDate, &movie.Rating, &actorID, &name, &gender, &birthdate)
		if err != nil {
			logger.Error("Ошибка сканирования результата:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		movie.ReleaseDate = releaseDate.Format("2006-01-02")
		movie.Actors = []models.Actor{{ID: actorID, Name: name, Gender: gender, Birthdate: birthdate.Format("2006-01-02")}}
		movies = append(movies, movie)
	}
	if err := rows.Err(); err != nil {
		logger.Error("Ошибка перебора результатов:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
