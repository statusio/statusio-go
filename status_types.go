package statusio

import "time"

type StatusSummaryResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result struct {
		Status []struct {
			ID         string `json:"id"`
			Name       string `json:"name"`
			Containers []struct {
				ID         string    `json:"id"`
				Name       string    `json:"name"`
				Updated    time.Time `json:"updated"`
				Status     string    `json:"status"`
				StatusCode int       `json:"status_code"`
			} `json:"containers"`
		} `json:"status"`
	} `json:"result"`
}
