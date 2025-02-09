package repositories

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/meybili19/edit-reservation-microservice/config"
)

type ReservationQueryResponse struct {
	Data struct {
		GetReservationById struct {
			ID           int     `json:"id"`
			UserID       int     `json:"userId"`
			CarID        int     `json:"vehicleId"`
			ParkingLotID int     `json:"parkingLotId"`
			StartDate    string  `json:"startDate"`
			EndDate      string  `json:"endDate"`
			Status       string  `json:"status"`
			TotalAmount  float64 `json:"totalAmount"`
		} `json:"getReservationById"`
	} `json:"data"`
}

func GetReservationByID(id int) (*ReservationQueryResponse, error) {
	url := config.GetQueryReservationURL()

	if url == "" {
		return nil, fmt.Errorf("missing QUERY_RESERVATION_URL in .env file")
	}

	query := fmt.Sprintf(`{"query": "query { getReservationById(id: %d) { id userId vehicleId parkingLotId startDate endDate status totalAmount } }"}`, id)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(query)))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request to GraphQL service: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("GraphQL service returned status: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var result ReservationQueryResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error unmarshaling GraphQL response: %w", err)
	}

	return &result, nil
}
