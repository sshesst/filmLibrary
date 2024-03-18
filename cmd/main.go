//package main
//
//import (
//	"filmLibrary/internal/routes"
//	database "filmLibrary/internal/storage"
//	"fmt"
//	"github.com/gorilla/handlers"
//	"log"
//	"net/http"
//)
//
//func main() {
//	configPath := "config"
//	pool, err := database.NewDBPool(configPath)
//	if err != nil {
//		log.Fatal("Error creating DB pool:", err)
//	}
//	defer pool.Close()
//
//	err = database.CreateTables(pool)
//	if err != nil {
//		log.Fatal("Error creating tables:", err)
//	}
//
//	fmt.Println("Success!")
//
//	routes.SetupRoutes()
//
//	port := ":8080"
//	fmt.Printf("Server is running on port %s\n", port)
//	http.ListenAndServe(port, handlers.CORS(
//		handlers.AllowedOrigins([]string{"*"}),
//		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
//		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
//	)(http.DefaultServeMux))
//}

package main

import (
	"filmLibrary/internal/routes"
	database "filmLibrary/internal/storage"
	"filmLibrary/pkg/logging"
	"github.com/gorilla/handlers"
	"net/http"
)

func main() {
	configPath := "config"

	logger := logging.NewLogger()

	pool, err := database.NewDBPool(configPath)
	if err != nil {
		logger.Error("Error creating DB pool:", err)
		return
	}
	defer pool.Close()

	err = database.CreateTables(pool)
	if err != nil {
		logger.Error("Error creating tables:", err)
		return
	}

	logger.Info("Success!")

	routes.SetupRoutes(*logger)

	port := ":8080"
	logger.Info("Server is running on port", port)
	http.ListenAndServe(port, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)(http.DefaultServeMux))
}
