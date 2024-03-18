package main

import (
	"filmLibrary/internal/routes"
	database "filmLibrary/internal/storage"
	"fmt"
	"github.com/gorilla/handlers"
	"log"
	"net/http"
)

func main() {
	configPath := "config"
	pool, err := database.NewDBPool(configPath)
	if err != nil {
		log.Fatal("Error creating DB pool:", err)
	}
	defer pool.Close()

	err = database.CreateTables(pool)
	if err != nil {
		log.Fatal("Error creating tables:", err)
	}

	fmt.Println("Success!")

	routes.SetupRoutes()

	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	http.ListenAndServe(port, handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	)(http.DefaultServeMux))
}
