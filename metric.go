package statusio

import "errors"

func (a StatusioApi) MetricUpdate(metric Metric) (r MetricUpdateResponse, err error) {
	err = a.apiRequest("POST", "metric/update", metric, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}
