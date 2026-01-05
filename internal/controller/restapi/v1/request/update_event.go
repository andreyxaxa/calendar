package request

import "github.com/andreyxaxa/calendar/pkg/types/date"

type UpdateRequest struct {
	UserID   int        `json:"user_id"`
	EventUID string     `json:"uid"`
	Date     *date.Date `json:"date"`
	Text     string     `json:"text"`
}
