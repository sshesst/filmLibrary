package routes

import (
	"filmLibrary/internal/controllers"
	"filmLibrary/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes() {
	r := mux.NewRouter()

	r.HandleFunc("/addActor", utils.BasicAuthMiddleware(adminOnlyMiddleware(controllers.AddActor))).Methods("POST")
	r.HandleFunc("/updateActor", utils.BasicAuthMiddleware(controllers.UpdateActor)).Methods("PUT")
	r.HandleFunc("/deleteActor/{id}", utils.BasicAuthMiddleware(controllers.DeleteActor)).Methods("DELETE")

	r.HandleFunc("/addMovie", utils.BasicAuthMiddleware(controllers.AddMovie)).Methods("POST")
	r.HandleFunc("/updateMovie", utils.BasicAuthMiddleware(controllers.UpdateMovie)).Methods("PUT")
	r.HandleFunc("/deleteMovie/{id}", utils.BasicAuthMiddleware(controllers.DeleteMovie)).Methods("DELETE")

	r.HandleFunc("/movies", utils.BasicAuthMiddleware(controllers.GetMovies)).Methods("GET")

	r.HandleFunc("/search-movies", utils.BasicAuthMiddleware(controllers.SearchMoviesByName)).Methods("GET")
	r.HandleFunc("/search-moviesby", utils.BasicAuthMiddleware(controllers.SearchMoviesByActor)).Methods("GET")
	r.HandleFunc("/all", utils.BasicAuthMiddleware(controllers.GetAllMovies)).Methods("GET")

	http.Handle("/", r)
}

// adminOnlyMiddleware проверяет, является ли пользователь администратором
func adminOnlyMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		isAdmin, ok := r.Context().Value("UserRole").(bool)
		if !ok || isAdmin == false {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
		next(w, r)
	}
}
