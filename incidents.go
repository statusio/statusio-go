package statusio

import (
	"errors"
	"fmt"
)

type IncidentListActiveResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result []IncidentResponse `json:"result"`
}

func (a StatusioApi) IncidentListActive(statusPageID string, page int, startDate, endDate int64) (r IncidentListActiveResponse, err error) {
	queryParams := fmt.Sprintf("?page=%d", page)
	if startDate > 0 {
		queryParams += fmt.Sprintf("&start_date=%d", startDate)
	}
	if endDate > 0 {
		queryParams += fmt.Sprintf("&end_date=%d", endDate)
	}
	
	err = a.apiRequest("GET", fmt.Sprintf("incident/list/active/%s%s", statusPageID, queryParams), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) IncidentList(statusPageID string) (r IncidentListResponse, err error) {
	err = a.apiRequest("GET", fmt.Sprintf("incident/list/%s", statusPageID), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) IncidentListByID(statusPageID string) (r IncidentListByIDResponse, err error) {
	err = a.apiRequest("GET", fmt.Sprintf("incidents/%s", statusPageID), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) IncidentMessage(statusPageID, messageID string) (r IncidentMessageResponse, err error) {
	err = a.apiRequest("GET", fmt.Sprintf("incident/message/%s/%s", statusPageID, messageID), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) IncidentSingle(statusPageID, incidentID string) (r IncidentSingleResponse, err error) {
	err = a.apiRequest("GET", fmt.Sprintf("incident/%s/%s", statusPageID, incidentID), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) IncidentCreate(incident Incident) (r IncidentCreateResponse, err error) {
	err = a.apiRequest("POST", "incident/create", incident, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) IncidentUpdate(incident Incident) (r IncidentUpdateResponse, err error) {
	err = a.apiRequest("POST", "incident/update", incident, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) IncidentResolve(incident Incident) (r IncidentResolveResponse, err error) {
	err = a.apiRequest("POST", "incident/resolve", incident, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) IncidentDelete(incident Incident) (r IncidentDeleteResponse, err error) {
	err = a.apiRequest("POST", "incident/delete", incident, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

