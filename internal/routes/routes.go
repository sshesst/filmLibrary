package routes

import (
	"filmLibrary/internal/controllers"
	"filmLibrary/internal/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes() {
	r := mux.NewRouter()

	r.HandleFunc("/actor/add-actor", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(controllers.AddActor))).Methods("POST")
	r.HandleFunc("/actor/update-actor", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(controllers.UpdateActor))).Methods("PUT")
	r.HandleFunc("/actor/delete-actor/{id}", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(controllers.DeleteActor))).Methods("DELETE")
	r.HandleFunc("/movie/add-movie", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(controllers.AddMovie))).Methods("POST")
	r.HandleFunc("/movie/update-movie", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(controllers.UpdateMovie))).Methods("PUT")
	r.HandleFunc("/movie/delete-movie/{id}", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(controllers.DeleteMovie))).Methods("DELETE")

	r.HandleFunc("/sort-movies", middleware.BasicAuthMiddleware(controllers.GetMovies)).Methods("GET")
	r.HandleFunc("/search-movies-by-title", middleware.BasicAuthMiddleware(controllers.SearchMoviesByName)).Methods("GET")
	r.HandleFunc("/search-movies-by-actor", middleware.BasicAuthMiddleware(controllers.SearchMoviesByActor)).Methods("GET")
	r.HandleFunc("/actor-list", middleware.BasicAuthMiddleware(controllers.GetAllMovies)).Methods("GET")

	http.Handle("/", r)
}
