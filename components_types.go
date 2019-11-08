package statusio

type ComponentStatus struct {
	StatuspageID  string `json:"statuspage_id"`
	Component     string `json:"component"`
	Container     string `json:"container"`
	Details       string `json:"details"`
	CurrentStatus Status `json:"current_status"`
}

type ComponentListResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result []struct {
		ID         string `json:"_id"`
		Statuspage string `json:"statuspage"`
		HookKey    string `json:"hook_key"`
		Containers []struct {
			ID          string `json:"_id"`
			Name        string `json:"name"`
			V           int    `json:"__v"`
			LocationGeo struct {
				QueryType   string   `json:"query_type"`
				Name        string   `json:"name"`
				Description string   `json:"description"`
				Region      string   `json:"region"`
				Country     string   `json:"country"`
				Address     string   `json:"address"`
				Host        string   `json:"host"`
				Coords      []string `json:"coords"`
			} `json:"location_geo"`
			Location string `json:"location"`
		} `json:"containers"`
		History []struct {
			MessageId string `json:"message_id"`
			Id        string `json:"_id"`
			Datetime  string `json:"datetime"`
		} `json:"history"`
		Name        string `json:"name"`
		V           int    `json:"__v"`
		Position    int    `json:"position"`
		Description string `json:"description"`
	} `json:"result"`
}

type ComponentUpdateResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result bool `json:"result"`
}
