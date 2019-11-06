package statusio

import "time"

type StatusSummaryResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result struct {
		Status []struct {
			ID         string    `json:"id"`
			Name       string    `json:"name"`
			Updated    time.Time `json:"updated"`
			Status     string    `json:"status"`
			StatusCode int       `json:"status_code"`
			Containers []struct {
				ID         string    `json:"id"`
				Name       string    `json:"name"`
				Updated    time.Time `json:"updated"`
				Status     string    `json:"status"`
				StatusCode int       `json:"status_code"`
			} `json:"containers"`
		} `json:"status"`
		StatusOverall struct {
			Updated    time.Time `json:"updated"`
			Status     string    `json:"status"`
			StatusCode int       `json:"status_code"`
		} `json:"status_overall"`
		Incidents []struct {
			ID           string    `json:"_id"`
			Name         string    `json:"name"`
			DateTimeOpen time.Time `json:"datetime_open"`
			Status       string    `json:"status"`
			StatusCode   int       `json:"status_code"`
			Messages     []struct {
				Details  string    `json:"details"`
				DateTime time.Time `json:"datetime"`
				State    int       `json:"state"`
				Status   int       `json:"status"`
			} `json:"messages"`
			Components []struct {
				ID   string `json:"_id"`
				Name string `json:"name"`
			} `json:"components_affected"`
			Containers []struct {
				ID   string `json:"_id"`
				Name string `json:"name"`
			} `json:"containers_affected"`
		} `json:"incidents"`
		Maintenance struct {
			Active []struct {
				ID            string    `json:"_id"`
				Name          string    `json:"name"`
				DateTimeOpen  time.Time `json:"datetime_open"`
				DateTimeStart time.Time `json:"datetime_planned_start"`
				DateTimeEnd   time.Time `json:"datetime_planned_end"`
				Messages      []struct {
					Details  string    `json:"details"`
					DateTime time.Time `json:"datetime"`
					State    int       `json:"state"`
					Status   int       `json:"status"`
				} `json:"messages"`
				Components []struct {
					ID   string `json:"_id"`
					Name string `json:"name"`
				} `json:"components_affected"`
				Containers []struct {
					ID   string `json:"_id"`
					Name string `json:"name"`
				} `json:"containers_affected"`
			} `json:"active"`
			Upcoming []struct {
				ID            string    `json:"_id"`
				Name          string    `json:"name"`
				DateTimeOpen  time.Time `json:"datetime_open"`
				DateTimeStart time.Time `json:"datetime_planned_start"`
				DateTimeEnd   time.Time `json:"datetime_planned_end"`
				Messages      []struct {
					Details  string    `json:"details"`
					DateTime time.Time `json:"datetime"`
					State    int       `json:"state"`
					Status   int       `json:"status"`
				} `json:"messages"`
				Components []struct {
					ID   string `json:"_id"`
					Name string `json:"name"`
				} `json:"components_affected"`
				Containers []struct {
					ID   string `json:"_id"`
					Name string `json:"name"`
				} `json:"containers_affected"`
			} `json:"upcoming"`
		} `json:"maintenance"`
	} `json:"result"`
}
