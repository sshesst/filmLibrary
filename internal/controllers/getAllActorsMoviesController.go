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

type MovieWithoutActors struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ReleaseDate string  `json:"releaseDate"`
	Rating      float64 `json:"rating"`
}

type ActorWithMovies struct {
	models.Actor
	Movies []MovieWithoutActors `json:"movies"`
}

func GetAllActorMovies(w http.ResponseWriter, r *http.Request, logger logging.Logger) {
	sqlQuery := `
        SELECT a.id, a.name, a.gender, a.birthdate, m.id, m.title, m.description, m.release_date, m.rating
        FROM actors a 
        JOIN movie_actors ma ON a.id = ma.actor_id 
        JOIN movies m ON ma.movie_id = m.id
    `

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

	actorsMap := make(map[uint]*ActorWithMovies)
	for rows.Next() {
		var actorID uint
		var actorName, actorGender string
		var actorBirthdate time.Time
		var movieID uint
		var movieTitle, movieDescription string
		var movieReleaseDate time.Time
		var movieRating float64

		err := rows.Scan(&actorID, &actorName, &actorGender, &actorBirthdate, &movieID, &movieTitle, &movieDescription, &movieReleaseDate, &movieRating)
		if err != nil {
			logger.Error("Ошибка сканирования результата:", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if actor, ok := actorsMap[actorID]; ok {
			actor.Movies = append(actor.Movies, MovieWithoutActors{
				ID:          movieID,
				Title:       movieTitle,
				Description: movieDescription,
				ReleaseDate: movieReleaseDate.Format("2006-01-02"),
				Rating:      movieRating,
			})
		} else {
			actorsMap[actorID] = &ActorWithMovies{
				Actor: models.Actor{
					ID:        actorID,
					Name:      actorName,
					Gender:    actorGender,
					Birthdate: actorBirthdate.Format("2006-01-02"),
				},
				Movies: []MovieWithoutActors{{
					ID:          movieID,
					Title:       movieTitle,
					Description: movieDescription,
					ReleaseDate: movieReleaseDate.Format("2006-01-02"),
					Rating:      movieRating,
				}},
			}
		}
	}
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var actorsWithMovies []ActorWithMovies
	for _, actor := range actorsMap {
		logger.Error("Ошибка перебора результатов:", err)
		actorsWithMovies = append(actorsWithMovies, *actor)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(actorsWithMovies)
}
