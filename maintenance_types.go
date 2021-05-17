package statusio

import "time"

type Maintenance struct {
	StatuspageID              string   `json:"statuspage_id"`
	AllInfrastructureAffected string   `json:"all_infrastructure_affected,omitempty"`
	InfrastructureAffected    []string `json:"infrastructure_affected"`
	Automation                string   `json:"automation"`
	MaintenanceName           string   `json:"maintenance_name"`
	MaintenanceDetails        string   `json:"maintenance_details"`
	MessageSubject            string   `json:"message_subject,omitempty"`
	DatePlannedStart          string   `json:"date_planned_start"`
	TimePlannedStart          string   `json:"time_planned_start"`
	DatePlannedEnd            string   `json:"date_planned_end"`
	TimePlannedEnd            string   `json:"time_planned_end"`
	MaintenanceNotifyNow      string   `json:"maintenance_notify_now"`
	MaintenanceNotify72Hr     string   `json:"maintenance_notify_72_hr"`
	MaintenanceNotify24Hr     string   `json:"maintenance_notify_24_hr"`
	MaintenanceNotify1Hr      string   `json:"maintenance_notify_1_hr"`
}

type MaintenanceControl struct {
	StatuspageID       string `json:"statuspage_id"`
	MaintenanceID      string `json:"maintenance_id"`
	MaintenanceDetails string `json:"maintenance_details,omitempty"`
	MessageSubject     string   `json:"message_subject,omitempty"`
	NotifyEmail        string `json:"notify_email"`
	NotifySms          string `json:"notify_sms"`
	NotifyWebhook      string `json:"notify_webhook"`
	Social             string `json:"social"`
	Irc                string `json:"irc,omitempty"`
	Hipchat            string `json:"hipchat,omitempty"`
	Msteams            string `json:"msteams,omitempty"`
	Slack              string `json:"slack,omitempty"`
}

type MaintenanceResponse struct {
	V                  int    `json:"__v"`
	ID                 string `json:"_id"`
	ComponentsAffected []struct {
		V          int      `json:"__v"`
		ID         string   `json:"_id"`
		Position   int      `json:"position"`
		Statuspage string   `json:"statuspage"`
		Containers []string `json:"containers"`
		Name       string   `json:"name"`
	} `json:"components_affected"`
	ContainersAffected []struct {
		V        int    `json:"__v"`
		ID       string `json:"_id"`
		Name     string `json:"name"`
		Location string `json:"location"`
	} `json:"containers_affected"`
	DatetimeOpen         time.Time `json:"datetime_open"`
	DatetimePlannedEnd   time.Time `json:"datetime_planned_end"`
	DatetimePlannedStart time.Time `json:"datetime_planned_start"`
	Kind                 string    `json:"kind"`
	Messages             []struct {
		Details     string    `json:"details"`
		Source      string    `json:"source"`
		State       int       `json:"state"`
		Status      int       `json:"status"`
		Statuspage  string    `json:"statuspage"`
		Maintenance string    `json:"maintenance"`
		IP          string    `json:"ip"`
		ID          string    `json:"_id"`
		V           int       `json:"__v"`
		Datetime    time.Time `json:"datetime"`
		Containers  []string  `json:"containers"`
		Components  []string  `json:"components"`
	} `json:"messages"`
	Name       string `json:"name"`
	Statuspage string `json:"statuspage"`
}

type MaintenanceListResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result struct {
		ActiveMaintenances   []MaintenanceResponse `json:"active_maintenances"`
		UpcomingMaintenances []MaintenanceResponse `json:"upcoming_maintenances"`
		ResolvedMaintenances []MaintenanceResponse `json:"resolved_maintenances"`
	} `json:"result"`
}

type MaintenanceListByIDResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result struct {
		ActiveMaintenances   []string `json:"active_maintenances"`
		UpcomingMaintenances []string `json:"upcoming_maintenances"`
		ResolvedMaintenances []string `json:"resolved_maintenances"`
	} `json:"result"`
}

type MaintenanceSingleResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result MaintenanceResponse `json:"result"`
}

type MaintenanceMessageResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result struct {
		Details     string    `json:"details"`
		Source      string    `json:"source"`
		State       int       `json:"state"`
		Statuspage  string    `json:"statuspage"`
		Maintenance string    `json:"maintenance"`
		IP          string    `json:"ip"`
		ID          string    `json:"_id"`
		V           int       `json:"__v"`
		Datetime    time.Time `json:"datetime"`
		Containers  []string  `json:"containers"`
		Components  []string  `json:"components"`
	} `json:"result"`
}

type MaintenanceScheduleResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result string `json:"result"`
}

type MaintenanceStartResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result bool `json:"result"`
}

type MaintenanceUpdateResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result bool `json:"result"`
}

type MaintenanceFinishResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result bool `json:"result"`
}

type MaintenanceDeleteResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result bool `json:"result"`
}
