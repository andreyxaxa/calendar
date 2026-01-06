package request

import "github.com/andreyxaxa/calendar/pkg/types/date"

// CreateRequest -.
type CreateRequest struct {
	UserID int        `json:"user_id"`
	Date   *date.Date `json:"date"`
	Text   string     `json:"text"`
}
