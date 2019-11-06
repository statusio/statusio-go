package statusio

import (
	"testing"
	"time"
)

const (
	apiID        = ""
	apiKey       = ""
	statusPageID = "568d8a3e3cada8c2490000dd"
	metricID     = "568d8ab5efe35d412f0006f8"
	component    = "568d8a3e3cada8c2490000ed"
	container    = "568d8a3e3cada8c2490000ec"
)

var componentContainerCombination = []string{"568d8a3e3cada8c2490000ed-568d8a3e3cada8c2490000ec"}

var (
	api *StatusioApi
	id1 string
	id2 string
)

func init() {
	api = NewStatusioApi(apiID, apiKey)
}

func Test_StatusioApi_StatusSummary(t *testing.T) {
	_, err := api.StatusSummary(statusPageID)

	if err != nil {
		t.Fatal(err)
	}
}

func Test_StatusioApi_SubscribersAdd(t *testing.T) {
	result, err := api.SubscriberAdd(Subscriber{
		StatuspageID: statusPageID,
		Method:       "email",
		Address:      "test@example.com",
		Granular:     "",
	})

	if err != nil {
		t.Fatal(err)
	}

	id1 = result.SubscriberID
}

func Test_StatusioApi_SubscribersUpdate(t *testing.T) {
	_, err := api.SubscriberUpdate(Subscriber{
		StatuspageID: statusPageID,
		SubscriberID: id1,
		Method:       "email",
		Address:      "test@example.com",
		Granular:     "",
	})

	if err != nil {
		t.Fatal(err)
	}
}

func Test_StatusioApi_SubscribersList(t *testing.T) {
	result, err := api.SubscriberList(statusPageID)

	if err != nil {
		t.Fatal(err)
	}

	if result.Result.Email[0].ID != id1 {
		t.Fatalf("First user ID is %s and stored id is %s", result.Result.Email[0].ID, id1)
	}
}

func Test_StatusioApi_SubscribersRemove(t *testing.T) {
	_, err := api.SubscriberRemove(Subscriber{
		StatuspageID: statusPageID,
		SubscriberID: id1,
	})

	if err != nil {
		t.Fatal(err)
	}
}

func Test_StatusioApi_Metric_Update(t *testing.T) {
	yesterday := time.Now().Add(-24 * time.Hour)
	yesterday = time.Date(yesterday.Year(), yesterday.Month(), yesterday.Day(), 0, 0, 0, 0, yesterday.Location())

	metric := Metric{
		StatuspageID: statusPageID,
		MetricID:     metricID,
		DayAvg:       22.58,
		DayStart:     yesterday.Unix() * 1000,
		DayDates:     []time.Time{},
		DayValues:    []float64{},
		WeekAvg:      20.07,
		WeekStart:    yesterday.Add(-7*24*time.Hour).Unix() * 1000,
		WeekDates:    []time.Time{},
		WeekValues:   []float64{},
		MonthAvg:     10.63,
		MonthStart:   yesterday.Add(-30*24*time.Hour).Unix() * 1000,
		MonthDates:   []time.Time{},
		MonthValues:  []float64{},
	}

	for hour := 0; hour < 24; hour++ {
		metric.DayDates = append(metric.DayDates, yesterday.Add(time.Duration(hour)*time.Hour))
		metric.DayValues = append(metric.DayValues, 10.0+(0.2*float64(hour)))
	}

	for wday := 0; wday < 7; wday++ {
		metric.WeekDates = append(metric.WeekDates, yesterday.Add(-7*24*time.Hour).Add(time.Duration(wday)*24*time.Hour))
		metric.WeekValues = append(metric.WeekValues, 10.0+(0.4*float64(wday)))
	}

	for mday := 0; mday < 30; mday++ {
		metric.MonthDates = append(metric.MonthDates, yesterday.Add(-30*24*time.Hour).Add(time.Duration(mday)*24*time.Hour))
		metric.MonthValues = append(metric.MonthValues, 10.0+(0.3*float64(mday)))
	}

	_, err := api.MetricUpdate(metric)

	if err != nil {
		t.Fatal(err)
	}
}

func Test_StatusioApi_Maintenance_Schedule(t *testing.T) {
	result, err := api.MaintenanceSchedule(Maintenance{
		StatuspageID:              statusPageID,
		InfrastructureAffected:    componentContainerCombination,
		MaintenanceName:           "Automated Test",
		MaintenanceDetails:        "Automated Test Details",
		DatePlannedStart:          time.Now().Add(1 * time.Hour).Format("2020/01/02"),
		TimePlannedStart:          time.Now().Add(1 * time.Hour).Format("15:04"),
		DatePlannedEnd:            time.Now().Add(2 * time.Hour).Format("2020/01/02"),
		TimePlannedEnd:            time.Now().Add(2 * time.Hour).Format("15:04"),
		MaintenanceNotifyNow:      "1",
		MaintenanceNotify1Hr:      "1",
		MaintenanceNotify24Hr:     "1",
		MaintenanceNotify72Hr:     "1",
		Automation:                "1",
		AllInfrastructureAffected: "1",
	})

	if err != nil {
		t.Fatal(err)
	}

	id1 = result.Result
}

func Test_StatusioApi_Maintenance_List(t *testing.T) {
	result, err := api.MaintenanceList(statusPageID)

	if err != nil {
		t.Fatal(err)
	}

	if len(result.Result.UpcomingMaintenances) == 0 {
		t.Fatal("Upcoming events list is empty!")
	}

	if result.Result.UpcomingMaintenances[0].ID != id1 {
		t.Fatalf("Upcoming event ID is %s and stored ID is %s", result.Result.UpcomingMaintenances[0].ID, id1)
	}

	id2 = result.Result.UpcomingMaintenances[0].Messages[0].ID
}

func Test_StatusioApi_Maintenance_Message(t *testing.T) {
	_, err := api.MaintenanceMessage(statusPageID, id2)

	if err != nil {
		t.Fatal(err)
	}
}

func Test_StatusioApi_Maintenance_Start(t *testing.T) {
	_, err := api.MaintenanceStart(MaintenanceControl{
		StatuspageID:       statusPageID,
		MaintenanceID:      id1,
		MaintenanceDetails: "Automated Test MaintenanceDetails START",
		NotifyEmail:        "0",
		NotifySms:          "0",
		NotifyWebhook:      "0",
		Social:             "0",
	})

	if err != nil {
		t.Fatal(err)
	}
}

func Test_StatusioApi_Maintenance_Update(t *testing.T) {
	_, err := api.MaintenanceUpdate(MaintenanceControl{
		StatuspageID:       statusPageID,
		MaintenanceID:      id1,
		MaintenanceDetails: "Automated Test MaintenanceDetails UPDATE",
		NotifyEmail:        "0",
		NotifySms:          "0",
		NotifyWebhook:      "0",
		Social:             "0",
	})

	if err != nil {
		t.Fatal(err)
	}
}

func Test_StatusioApi_Maintenance_Finish(t *testing.T) {
	_, err := api.MaintenanceFinish(MaintenanceControl{
		StatuspageID:       statusPageID,
		MaintenanceID:      id1,
		MaintenanceDetails: "Automated Test MaintenanceDetails FINISH",
		NotifyEmail:        "0",
		NotifySms:          "0",
		NotifyWebhook:      "0",
		Social:             "0",
	})

	if err != nil {
		t.Fatal(err)
	}
}

func Test_StatusioApi_Maintenance_Delete(t *testing.T) {
	_, err := api.MaintenanceDelete(MaintenanceControl{
		StatuspageID:  statusPageID,
		MaintenanceID: id1,
	})

	if err != nil {
		t.Fatal(err)
	}
}

func Test_StatusioApi_Incident_Create(t *testing.T) {
	result, err := api.IncidentCreate(Incident{
		StatuspageID:              statusPageID,
		AllInfrastructureAffected: "1",
		InfrastructureAffected:    componentContainerCombination,
		IncidentName:              "Automated Test",
		IncidentDetails:           "Automated Test Details",
		CurrentState:              STATE_INVESTIGATING,
		CurrentStatus:             STATUS_OPERATIONAL,
		NotifyEmail:               "0",
		NotifySms:                 "0",
		NotifyWebhook:             "0",
		Social:                    "0",
	})

	if err != nil {
		t.Fatal(err)
	}

	id1 = result.Result
}

func Test_StatusioApi_Incident_List(t *testing.T) {
	result, err := api.IncidentList(statusPageID)

	if err != nil {
		t.Fatal(err)
	}

	if len(result.Result.ActiveIncidents) == 0 {
		t.Fatal("Active incidents list is empty!")
	}

	if result.Result.ActiveIncidents[0].ID != id1 {
		t.Fatalf("Active incident ID is %s and stored ID is %s", result.Result.ActiveIncidents[0].ID, id1)
	}

	id2 = result.Result.ActiveIncidents[0].Messages[0].ID
}

func Test_StatusioApi_Incident_Message(t *testing.T) {
	_, err := api.IncidentMessage(statusPageID, id2)

	if err != nil {
		t.Fatal(err)
	}
}

func Test_StatusioApi_Incident_Update(t *testing.T) {
	_, err := api.IncidentUpdate(Incident{
		StatuspageID:    statusPageID,
		IncidentID:      id1,
		IncidentDetails: "Automated Test Details Updated",
		CurrentState:    STATE_IDENTIFIED,
		CurrentStatus:   STATUS_OPERATIONAL,
		NotifyEmail:     "0",
		NotifySms:       "0",
		NotifyWebhook:   "0",
		Social:          "0",
	})

	if err != nil {
		t.Fatal(err)
	}
}

func Test_StatusioApi_Incident_Resolve(t *testing.T) {
	_, err := api.IncidentResolve(Incident{
		StatuspageID:    statusPageID,
		IncidentID:      id1,
		IncidentDetails: "Automated Test Details Resolved",
		NotifyEmail:     "0",
		NotifySms:       "0",
		NotifyWebhook:   "0",
		Social:          "0",
	})

	if err != nil {
		t.Fatal(err)
	}
}

func Test_StatusioApi_Incident_Delete(t *testing.T) {
	_, err := api.IncidentDelete(Incident{
		StatuspageID: statusPageID,
		IncidentID:   id1,
	})

	if err != nil {
		t.Fatal(err)
	}
}

func Test_StatusioApi_Component_List(t *testing.T) {
	_, err := api.ComponentList(statusPageID)

	if err != nil {
		t.Fatal(err)
	}
}

func Test_StatusioApi_Component_Update(t *testing.T) {
	_, err := api.ComponentUpdate(ComponentStatus{
		StatuspageID:  statusPageID,
		Component:     component,
		Container:     container,
		Details:       "Automated Test Update",
		CurrentStatus: STATUS_DEGRADED_PERFORMANCE,
	})

	if err != nil {
		t.Fatal(err)
	}
}
