//Package statusio provides structs and functions for accessing version 2.0
//of the Status.io API.
//
//Successful API queries return native Go structs that can be used immediately,
//with no need for type assertions.
//
//Authentication
//
//If you already have the API ID and API KEY for your user, creating the client is simple:
//
//  api := statusio.NewStatusioApi("your-api-id", "your-api-key")
//
//
//Queries
//
//Executing queries is simple
//
//  statusSummary, _ := api.StatusSummary("your-statuspage-id")
//  fmt.Print(result.Result.Status[0].Containers[0].Status)
//
//Endpoints
//
//`statusio` implements most of the endpoints defined in the Status.io documentation: http://docs.statusio.apiary.io
//For clarity, in most cases, the function name is simply the name of the endpoint (e.g., the endpoint `status/summary` is provided by the function `StatusSummary`).
//
//More detailed information about the behavior of each particular endpoint can be found at the official Status.io documentation.
package statusio

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	baseURLV2 = "https://api.status.io/v2/"

	STATE_INVESTIGATING = 100 // Deprecated: use StateInvestigating instead
	STATE_IDENTIFIED    = 200 // Deprecated: use StateIdentified instead
	STATE_MONITORING    = 300 // Deprecated: use StateMonitoring instead

	STATUS_OPERATIONAL                = 100 // Deprecated: use StatusOperational instead
	STATUS_PLANNED_MAINTENANCE        = 200 // Deprecated: use StatusPlannedMaintenance instead
	STATUS_DEGRADED_PERFORMANCE       = 300 // Deprecated: use StatusDegradedPerformance instead
	STATUS_PARTIAL_SERVICE_DISRUPTION = 400 // Deprecated: use StatusPartialServiceDisruption instead
	STATUS_SERVICE_DISRUPTION         = 500 // Deprecated: use StatusServiceDisruption instead
	STATUS_SECURITY_EVENT             = 600 // Deprecated: use StatusSecurityEvent instead
)

type State int
type Status int

// State enum values
var (
	StateInvestigating State = 100
	StateIdentified    State = 200
	StateMonitoring    State = 300
)

// Status enum values
var (
	StatusOperational              Status = 100
	StatusPlannedMaintenance       Status = 200
	StatusDegradedPerformance      Status = 300
	StatusPartialServiceDisruption Status = 400
	StatusServiceDisruption        Status = 500
	StatusSecurityEvent            Status = 600
)

type StatusioApi struct {
	ApiId      string
	ApiKey     string
	HttpClient *http.Client
	baseUrl    string
}

func (a StatusioApi) apiRequest(method string, resource string, data interface{}, result interface{}) error {
	var (
		req *http.Request
		err error
	)

	if data != nil {
		b, marshalErr := json.Marshal(data)
		if marshalErr != nil {
			return marshalErr
		}
		req, err = http.NewRequest(method, fmt.Sprintf("%s%s", a.baseUrl, resource), bytes.NewReader(b))
	} else {
		req, err = http.NewRequest(method, fmt.Sprintf("%s%s", a.baseUrl, resource), nil)
	}

	if err != nil {
		return err
	}

	req.Header.Set("x-api-id", a.ApiId)
	req.Header.Set("x-api-key", a.ApiKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := a.HttpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("response status code is %d", resp.StatusCode)
	}
	return json.NewDecoder(resp.Body).Decode(result)
}

func NewStatusioApi(api_id string, api_key string) *StatusioApi {
	c := &StatusioApi{
		ApiId:      api_id,
		ApiKey:     api_key,
		HttpClient: http.DefaultClient,
		baseUrl:    baseURLV2,
	}
	return c
}
