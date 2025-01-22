package services

import (
	"database/sql"
	"fmt"

	"github.com/meybili19/edit-reservation-microservice/repositories"
)

func UpdateReservationService(db *sql.DB, reservation map[string]interface{}) error {
	id, ok := reservation["id"].(float64)
	if !ok {
		return fmt.Errorf("id must be a valid number")
	}
	reservation["id"] = int(id)

	return repositories.UpdateReservation(db, reservation)
}
