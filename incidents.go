package statusio

import (
	"errors"
	"fmt"
)

func (a StatusioApi) IncidentList(statuspage_id string) (r IncidentListResponse, err error) {
	err = a.apiRequest("GET", fmt.Sprintf("incident/list/%s", statuspage_id), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) IncidentListByID(statuspage_id string) (r IncidentListByIDResponse, err error) {
	err = a.apiRequest("GET", fmt.Sprintf("incidents/%s", statuspage_id), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) IncidentMessage(statuspage_id string, message_id string) (r IncidentMessageResponse, err error) {
	err = a.apiRequest("GET", fmt.Sprintf("incident/message/%s/%s", statuspage_id, message_id), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) IncidentSingle(statuspage_id string, incident_id string) (r IncidentSingleResponse, err error) {
	err = a.apiRequest("GET", fmt.Sprintf("incident/%s/%s", statuspage_id, incident_id), nil, &r)
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
