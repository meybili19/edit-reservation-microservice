package repositories

import (
	"database/sql"
	"fmt"
	"time"
)

func UpdateReservation(db *sql.DB, reservation map[string]interface{}) error {
	query := `UPDATE Reservations SET car_id = ?, parking_lot_id = ?, start_date = ?, end_date = ?, total_amount = ? WHERE id = ?`

	startDate, err := time.Parse(time.RFC3339, reservation["start_date"].(string))
	if err != nil {
		return fmt.Errorf("invalid start_date: %w", err)
	}
	endDate, err := time.Parse(time.RFC3339, reservation["end_date"].(string))
	if err != nil {
		return fmt.Errorf("invalid end_date: %w", err)
	}

	duration := endDate.Sub(startDate).Hours()
	totalAmount := duration * 10 // Ejemplo: costo por hora = 10

	_, err = db.Exec(query, reservation["car_id"], reservation["parking_lot_id"], startDate, endDate, totalAmount, reservation["id"])
	return err
}
