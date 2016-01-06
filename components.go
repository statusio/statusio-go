package statusio

import (
	"errors"
	"fmt"
)

func (a StatusioApi) ComponentList(statuspage_id string) (r ComponentListResponse, err error) {
	err = a.apiRequest("GET", fmt.Sprintf("component/list/%s", statuspage_id), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) ComponentUpdate(status ComponentStatus) (r ComponentUpdateResponse, err error) {
	err = a.apiRequest("POST", "component/status/update", status, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}
