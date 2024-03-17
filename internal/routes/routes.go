package routes

import (
	"filmLibrary/internal/controllers"
	"filmLibrary/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes() {
	r := mux.NewRouter()

	r.HandleFunc("/addActor", utils.BasicAuthMiddleware(controllers.AddActor)).Methods("POST")
	r.HandleFunc("/updateActor", controllers.UpdateActor).Methods("PUT")
	r.HandleFunc("/deleteActor/{id}", controllers.DeleteActor).Methods("DELETE")

	r.HandleFunc("/addMovie", controllers.AddMovie).Methods("POST")
	r.HandleFunc("/updateMovie", controllers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/deleteMovie/{id}", controllers.DeleteMovie).Methods("DELETE")

	r.HandleFunc("/movies", controllers.GetMovies).Methods("GET")

	r.HandleFunc("/search-movies", controllers.SearchMoviesByName).Methods("GET")
	r.HandleFunc("/search-moviesby", controllers.SearchMoviesByActor).Methods("GET")
	r.HandleFunc("/all", controllers.GetAllMovies).Methods("GET")

	http.Handle("/", r)
}
