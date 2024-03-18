package main

import (
	"filmLibrary/internal/routes"
	database "filmLibrary/internal/storage"
	"filmLibrary/pkg/logging"
	"net/http"
)

// @title Todo App API
// @version 1.0
// @description API Server for TodoList Application

// @host localhost:8000
// @BasePath /

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
	http.ListenAndServe(port, nil)
}
