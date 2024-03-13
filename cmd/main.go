package main

import (
	database "filmLibrary"
	"fmt"
	"log"
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
}
