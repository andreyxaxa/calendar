package request

type DeleteRequest struct {
	UserID   int    `json:"user_id"`
	EventUID string `json:"uid"`
}
