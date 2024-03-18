package routes

import (
	_ "filmLibrary/docs"
	"filmLibrary/internal/controllers"
	"filmLibrary/internal/middleware"
	"filmLibrary/pkg/logging"
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
)

func SetupRoutes(logger logging.Logger) {

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

	http.HandleFunc("/actor/add-actor", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(wrappedAddActor)))
	http.HandleFunc("/actor/update-actor", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(wrappedUpdateActor)))
	http.HandleFunc("/actor/delete-actor/{id}", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(wrappedDeleteActor)))

	http.HandleFunc("/movie/add-movie", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(wrappedAddMovie)))
	http.HandleFunc("/movie/update-movie", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(wrappedUpdateMovie)))
	http.HandleFunc("/movie/delete-movie/{id}", middleware.BasicAuthMiddleware(middleware.AdminOnlyMiddleware(wrappedDeleteMovie)))
	http.HandleFunc("/swagger/", httpSwagger.WrapHandler)
	http.HandleFunc("/sort-movies", middleware.BasicAuthMiddleware(wrappedGetMovies))
	http.HandleFunc("/search-movies-by-title", middleware.BasicAuthMiddleware(wrappedSearchMoviesByName))
	http.HandleFunc("/search-movies-by-actor", middleware.BasicAuthMiddleware(wrappedSearchMoviesByActor))
	http.HandleFunc("/actor-list", middleware.BasicAuthMiddleware(wrappedGetAllMovies))

}
