package statusio

import (
	"errors"
	"fmt"
)

func (a StatusioApi) MaintenanceList(statusPageID string) (r MaintenanceListResponse, err error) {
	err = a.apiRequest("GET", fmt.Sprintf("maintenance/list/%s", statusPageID), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) MaintenanceListByID(statusPageID string) (r MaintenanceListByIDResponse, err error) {
	err = a.apiRequest("GET", fmt.Sprintf("maintenances/%s", statusPageID), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) MaintenanceMessage(statusPageID, messageID string) (r MaintenanceMessageResponse, err error) {
	err = a.apiRequest("GET", fmt.Sprintf("maintenance/message/%s/%s", statusPageID, messageID), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) MaintenanceSingle(statusPageID, maintenanceID string) (r MaintenanceSingleResponse, err error) {
	err = a.apiRequest("GET", fmt.Sprintf("maintenance/%s/%s", statusPageID, maintenanceID), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) MaintenanceSchedule(task Maintenance) (r MaintenanceScheduleResponse, err error) {
	err = a.apiRequest("POST", "maintenance/schedule", task, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) MaintenanceStart(control MaintenanceControl) (r MaintenanceStartResponse, err error) {
	err = a.apiRequest("POST", "maintenance/start", control, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) MaintenanceUpdate(control MaintenanceControl) (r MaintenanceUpdateResponse, err error) {
	err = a.apiRequest("POST", "maintenance/update", control, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) MaintenanceFinish(control MaintenanceControl) (r MaintenanceFinishResponse, err error) {
	err = a.apiRequest("POST", "maintenance/finish", control, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) MaintenanceDelete(control MaintenanceControl) (r MaintenanceDeleteResponse, err error) {
	err = a.apiRequest("POST", "maintenance/delete", control, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}
