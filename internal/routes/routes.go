package routes

import (
	"filmLibrary/internal/controllers"
	"filmLibrary/internal/middleware"
	"filmLibrary/pkg/logging"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes(logger logging.Logger) {
	r := mux.NewRouter()

	wrappedAddActor := func(w http.ResponseWriter, r *http.Request) {
		controllers.AddActor(w, r, logger)
	}

	wrappedUpdateActor := func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateActor(w, r, logger)
	}

	wrappedDeleteActor := func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteActor(w, r, logger)
	}

	wrappedAddMovie := func(w http.ResponseWriter, r *http.Request) {
		controllers.AddMovie(w, r, logger)
	}

	wrappedUpdateMovie := func(w http.ResponseWriter, r *http.Request) {
		controllers.UpdateMovie(w, r, logger)
	}

	wrappedDeleteMovie := func(w http.ResponseWriter, r *http.Request) {
		controllers.DeleteMovie(w, r, logger)
	}

	wrappedGetMovies := func(w http.ResponseWriter, r *http.Request) {
		controllers.GetMovies(w, r, logger)
	}

	wrappedSearchMoviesByName := func(w http.ResponseWriter, r *http.Request) {
		controllers.SearchMoviesByName(w, r, logger)
	}

	wrappedSearchMoviesByActor := func(w http.ResponseWriter, r *http.Request) {
		controllers.SearchMoviesByActor(w, r, logger)
	}

	wrappedGetAllMovies := func(w http.ResponseWriter, r *http.Request) {
		controllers.GetAllActorMovies(w, r, logger)
	}

	r.HandleFunc("/actor/add-actor", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(wrappedAddActor))).Methods("POST")
	r.HandleFunc("/actor/update-actor", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(wrappedUpdateActor))).Methods("PUT")
	r.HandleFunc("/actor/delete-actor/{id}", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(wrappedDeleteActor))).Methods("DELETE")

	r.HandleFunc("/movie/add-movie", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(wrappedAddMovie))).Methods("POST")
	r.HandleFunc("/movie/update-movie", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(wrappedUpdateMovie))).Methods("PUT")
	r.HandleFunc("/movie/delete-movie/{id}", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(wrappedDeleteMovie))).Methods("DELETE")

	r.HandleFunc("/sort-movies", middleware.BasicAuthMiddleware(wrappedGetMovies)).Methods("GET")
	r.HandleFunc("/search-movies-by-title", middleware.BasicAuthMiddleware(wrappedSearchMoviesByName)).Methods("GET")
	r.HandleFunc("/search-movies-by-actor", middleware.BasicAuthMiddleware(wrappedSearchMoviesByActor)).Methods("GET")
	r.HandleFunc("/actor-list", middleware.BasicAuthMiddleware(wrappedGetAllMovies)).Methods("GET")

	http.Handle("/", r)
}
