package models

import (
	"time"
)

// Actor модель для таблицы актёров
type Actor struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Birthdate time.Time `json:"birthdate"`
}

// Movie модель для таблицы фильмов
type Movie struct {
	ID          uint      `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ReleaseDate time.Time `json:"releaseDate"`
	Rating      float64   `json:"rating" check:"(rating >= 0) AND (rating <= 10)"`
	Actors      []Actor   `json:"actors"`
}

// User модель для таблицы пользователей
//type User struct {
//	ID       uint   `json:"id"`
//	Username string `json:"username"`
//	Password string `json:"-"`
//	Role     string `json:"role"`
//}

// связь Many-to-Many для Movie и Actor
type MovieActor struct {
	MovieID uint `json:"movieId"`
	ActorID uint `json:"actorId"`
}
