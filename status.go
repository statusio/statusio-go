package statusio

import (
	"errors"
	"fmt"
)

func (a StatusioApi) StatusSummary(StatusPageId string) (r StatusSummaryResponse, err error) {
	err = a.apiRequest("GET", fmt.Sprintf("status/summary/%s", StatusPageId), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}
