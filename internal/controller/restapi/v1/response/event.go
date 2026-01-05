package response

import "github.com/andreyxaxa/calendar/pkg/types/date"

type Response struct {
	Result ResultEvent `json:"result"`
}

type ResultEvent struct {
	UID    string    `json:"uid"`
	UserID int       `json:"user_id"`
	Date   date.Date `json:"date"`
	Text   string    `json:"text"`
}
