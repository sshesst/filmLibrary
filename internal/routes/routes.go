package routes

// SetupRoutes настраивает маршруты для вашего приложения
//func SetupRoutes() {
//	r := mux.NewRouter()
//	// Включаем CORS
//	corsHandler := handlers.CORS(
//		handlers.AllowedOrigins([]string{"*"}),
//		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
//		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
//	)
//
//	// Добавление актера
//	r.HandleFunc("/addActor", controllers.AddActor).Methods("POST")
//
//	// Обновление информации об актере
//	r.HandleFunc("/updateActor", controllers.UpdateActor).Methods("PUT")
//
//	// Удаление актера
//	r.HandleFunc("/deleteActor", controllers.DeleteActor).Methods("DELETE")
//
//	// Задаем обработчик CORS для всех маршрутов
//	http.Handle("/", corsHandler(r))
//}

import (
	"filmLibrary/internal/controllers"
	"github.com/gorilla/mux"
	"net/http"
)

// SetupRoutes настраивает маршруты для вашего приложения
func SetupRoutes() {
	// Создаем новый маршрутизатор
	r := mux.NewRouter()

	// Задаем обработчики для каждого маршрута
	r.HandleFunc("/addActor", controllers.AddActor).Methods("POST")
	r.HandleFunc("/updateActor", controllers.UpdateActor).Methods("PUT")
	r.HandleFunc("/deleteActor/{id}", controllers.DeleteActor).Methods("DELETE")

	r.HandleFunc("/addMovie", controllers.AddMovie).Methods("POST")
	r.HandleFunc("/updateMovie", controllers.UpdateMovie).Methods("PUT")
	r.HandleFunc("/deleteMovie/{id}", controllers.DeleteMovie).Methods("DELETE")

	// Используем маршрутизатор в качестве основного маршрута для HTTP сервера
	http.Handle("/", r)
}
