package statusio

import "time"

type Incident struct {
	StatuspageID              string   `json:"statuspage_id"`
	IncidentID                string   `json:"incident_id,omitempty"`
	IncidentName              string   `json:"incident_name,omitempty"`
	IncidentDetails           string   `json:"incident_details"`
	NotifyEmail               int      `json:"notify_email"`
	NotifySms                 int      `json:"notify_sms"`
	NotifyWebhook             int      `json:"notify_webhook"`
	Social                    int      `json:"social"`
	Irc                       int      `json:"irc,omitempty"`
	Hipchat                   int      `json:"hipchat,omitempty"`
	Slack                     int      `json:"slack,omitempty"`
	CurrentStatus             int      `json:"current_status,omitempty"`
	CurrentState              int      `json:"current_state,omitempty"`
	AllInfrastructureAffected int      `json:"all_infrastructure_affected,omitempty"`
	InfrastructureAffected    []string `json:"infrastructure_affected"`
}

type IncidentResponse struct {
	V                  int    `json:"__v"`
	ID                 string `json:"_id"`
	ComponentsAffected []struct {
		V          int    `json:"__v"`
		ID         string `json:"_id"`
		Position   int    `json:"position"`
		Statuspage string `json:"statuspage"`
		Containers []string `json:"containers"`
		Name       string   `json:"name"`
	} `json:"components_affected"`
	ContainersAffected []struct {
		V        int    `json:"__v"`
		ID       string `json:"_id"`
		Name     string `json:"name"`
		Location string `json:"location"`
	} `json:"containers_affected"`
	DatetimeOpen time.Time `json:"datetime_open"`
	Kind         string    `json:"kind"`
	Messages     []struct {
		Details    string    `json:"details"`
		Source     string    `json:"source"`
		State      int       `json:"state"`
		Status     int       `json:"status"`
		Statuspage string    `json:"statuspage"`
		Incident   string    `json:"incident"`
		IP         string    `json:"ip"`
		ID         string    `json:"_id"`
		V          int       `json:"__v"`
		Datetime   time.Time `json:"datetime"`
		Containers []string  `json:"containers"`
		Components []string  `json:"components"`
	} `json:"messages"`
	Name       string `json:"name"`
	Statuspage string `json:"statuspage"`
}

type IncidentListResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result struct {
		ActiveIncidents   []IncidentResponse `json:"active_incidents"`
		ResolvedIncidents []IncidentResponse `json:"resolved_incidents"`
	} `json:"result"`
}

type IncidentListByIDResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result struct {
		ActiveIncidents   []string `json:"active_incidents"`
		ResolvedIncidents []string `json:"resolved_incidents"`
	} `json:"result"`
}

type IncidentMessageResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result struct {
		UserEmail    string    `json:"user_email"`
		UserFullName string    `json:"user_full_name"`
		Details      string    `json:"details"`
		Source       string    `json:"source"`
		Social       string    `json:"social"`
		State        int       `json:"state"`
		Status       int       `json:"status"`
		Statuspage   string    `json:"statuspage"`
		Incident     string    `json:"incident"`
		IP           string    `json:"ip"`
		ID           string    `json:"_id"`
		V            int       `json:"__v"`
		Datetime     time.Time `json:"datetime"`
		Containers   []string  `json:"containers"`
		Components   []string  `json:"components"`
	} `json:"result"`
}

type IncidentCreateResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result string `json:"result"`
}

type IncidentUpdateResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result bool `json:"result"`
}

type IncidentResolveResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result bool `json:"result"`
}

type IncidentDeleteResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result bool `json:"result"`
}
