package routes

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/meybili19/edit-reservation-microservice/services"
)

func UpdateReservationHandler(databases map[string]*sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		var reservation map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&reservation); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		err := services.UpdateReservationService(databases["reservations"], reservation)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Reservation updated successfully",
		})
	}
}
