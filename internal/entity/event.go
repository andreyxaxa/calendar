package entity

import "time"

// Event -.
type Event struct {
	Date time.Time `json:"date"`
	Text string    `json:"text"`
}
