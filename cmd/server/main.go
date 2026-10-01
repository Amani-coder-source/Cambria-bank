package main

import (
	"fmt"
	"log"
	"net/http"

	"cambria-bank/database"
	"cambria-bank/routes"
)

func main() {
	database.ConnectDB()

	router := routes.SetupRoutes()

	fmt.Println("Cambria Capital Bank server running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", router))
}
