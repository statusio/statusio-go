package statusio

import (
	"errors"
	"fmt"
)

func (a StatusioApi) SubscriberList(statuspage_id string) (r SubscriberListResponse, err error) {
	err = a.apiRequest("GET", fmt.Sprintf("subscriber/list/%s", statuspage_id), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) SubscriberAdd(subscriber Subscriber) (r SubscriberAddResponse, err error) {
	err = a.apiRequest("POST", "subscriber/add", subscriber, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) SubscriberUpdate(subscriber Subscriber) (r SubscriberUpdateResponse, err error) {
	err = a.apiRequest("PATCH", "subscriber/update", subscriber, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}

func (a StatusioApi) SubscriberRemove(subscriber Subscriber) (r SubscriberRemoveResponse, err error) {
	err = a.apiRequest("DELETE", fmt.Sprintf("subscriber/remove/%s/%s", subscriber.StatuspageID, subscriber.SubscriberID), nil, &r)
	if r.Status.Error != "no" {
		err = errors.New(r.Status.Message)
	}
	return r, err
}
