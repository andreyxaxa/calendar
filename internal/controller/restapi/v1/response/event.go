package response

import "github.com/andreyxaxa/calendar/pkg/types/date"

// Response -.
type Response struct {
	Result ResultEvent `json:"result"`
}

// ResultEvent -.
type ResultEvent struct {
	UserID int       `json:"user_id"`
	UID    string    `json:"uid"`
	Date   date.Date `json:"date"`
	Text   string    `json:"text"`
}
