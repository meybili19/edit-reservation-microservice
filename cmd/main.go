package main

import (
	"log"
	"net/http"

	"github.com/meybili19/edit-reservation-microservice/config"
	"github.com/meybili19/edit-reservation-microservice/routes"
)

func main() {
	databases, err := config.InitDatabases()
	if err != nil {
		log.Fatalf("Error initializing databases: %v", err)
	}
	defer func() {
		for _, db := range databases {
			db.Close()
		}
	}()
	log.Println("All databases connected successfully!")

	http.HandleFunc("/reservations/update", routes.UpdateReservationHandler(databases))
	log.Println("Server running on port 8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
