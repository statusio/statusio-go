package statusio

import (
	"time"
)

type Subscriber struct {
	StatuspageID string `json:"statuspage_id"`
	SubscriberID string `json:"subscriber_id,omitempty"`
	Method       string `json:"method"`
	Address      string `json:"address"`
	Silent       string `json:"silent,omitempty"`
	Granular     string `json:"granular,omitempty"`
}

type SubscriberResponse struct {
	ID         string    `json:"_id"`
	Address    string    `json:"address"`
	Method     string    `json:"method"`
	Statuspage string    `json:"statuspage"`
	V          int       `json:"__v"`
	Granular   []string  `json:"granular"`
	Active     bool      `json:"active"`
	JoinDate   time.Time `json:"join_date"`
}

type SubscriberListResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result struct {
		Email   []SubscriberResponse `json:"email"`
		Webhook []SubscriberResponse `json:"webhook"`
		Sms     []SubscriberResponse `json:"sms"`
	} `json:"result"`
}

type SubscriberAddResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result       bool   `json:"result"`
	SubscriberID string `json:"subscriber_id"`
}

type SubscriberUpdateResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result SubscriberResponse `json:"result"`
}

type SubscriberRemoveResponse struct {
	Status struct {
		Error   string `json:"error"`
		Message string `json:"message"`
	} `json:"status"`
	Result bool `json:"result"`
}
