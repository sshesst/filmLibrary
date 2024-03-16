package models

// Actor модель для таблицы актёров
type Actor struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthdate"`
}

// Movie модель для таблицы фильмов
type Movie struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ReleaseDate string  `json:"releaseDate"`
	Rating      float64 `json:"rating"`
	Actors      []Actor `json:"actors"`
}

// связь Many-to-Many для Movie и Actor
type MovieActor struct {
	MovieID uint `json:"movieId" pg:",pk"`
	ActorID uint `json:"actorId" pg:",pk"`
}

// User модель для таблицы пользователей
//type User struct {
//	ID       uint   `json:"id"`
//	Username string `json:"username"`
//	Password string `json:"-"`
//	Role     string `json:"role"`
//}
