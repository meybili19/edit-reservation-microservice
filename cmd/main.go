package main

import (
	"log"
	"net/http"

	"github.com/meybili19/edit-reservation-microservice/config"
	"github.com/meybili19/edit-reservation-microservice/routes"
	"github.com/rs/cors"
)

func main() {
	if err := config.LoadEnv(); err != nil {
		log.Fatalf("Error loading environment variables: %v", err)
	}

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

	mux := http.NewServeMux()
	mux.HandleFunc("/reservations/update", routes.UpdateReservationHandler(databases))

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Permitir solicitudes desde el frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := corsHandler.Handler(mux)

	log.Println("Server running on port 4001")
	log.Fatal(http.ListenAndServe(":4001", handler))
}
