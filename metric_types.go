package statusio

import "time"

type Metric struct {
	StatuspageID string      `json:"statuspage_id"`
	MetricID     string      `json:"metric_id"`
	DayAvg       float64     `json:"day_avg"`
	DayStart     int64       `json:"day_start"`
	DayDates     []time.Time `json:"day_dates"`
	DayValues    []float64   `json:"day_values"`
	WeekAvg      float64     `json:"week_avg"`
	WeekStart    int64       `json:"week_start"`
	WeekDates    []time.Time `json:"week_dates"`
	WeekValues   []float64   `json:"week_values"`
	MonthAvg     float64     `json:"month_avg"`
	MonthStart   int64       `json:"month_start"`
	MonthDates   []time.Time `json:"month_dates"`
	MonthValues  []float64   `json:"month_values"`
}

type MetricUpdateResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result bool `json:"result"`
}
