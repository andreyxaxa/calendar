package entity

import "time"

type Event struct {
	Date time.Time `json:"date"`
	Text string    `json:"text"`
}
