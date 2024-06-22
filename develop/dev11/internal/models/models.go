package models

import (
	"time"
)

type Event struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Date       time.Time `json:"date"`
	Descrtpion string    `json:"descrtpion"`
}

type SuccessfulRequest struct {
	Result Event `json:"result"`
}

type SuccessfulRequestGet struct {
	Results []Event `json:"result"`
}
