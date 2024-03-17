package models

type Actor struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Gender    string `json:"gender"`
	Birthdate string `json:"birthdate"`
}

type Movie struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	ReleaseDate string  `json:"releaseDate"`
	Rating      float64 `json:"rating"`
	Actors      []Actor `json:"actors"`
}

type MovieActor struct {
	MovieID uint `json:"movieId" pg:",pk"`
	ActorID uint `json:"actorId" pg:",pk"`
}

type User struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"isadmin"`
	IsAuth   bool   `json:"isAuth"`
}
