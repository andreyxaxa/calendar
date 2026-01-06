package v1

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/andreyxaxa/calendar/internal/controller/restapi/v1/request"
	"github.com/andreyxaxa/calendar/internal/controller/restapi/v1/response"
	"github.com/andreyxaxa/calendar/internal/entity"
	"github.com/andreyxaxa/calendar/pkg/types/date"
	"github.com/andreyxaxa/calendar/pkg/types/errs"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// @Summary Create
// @Description Creates event by user_id, date, text
// @ID create
// @Tags events
// @Accept json
// @Produce json
// @Param request body request.CreateRequest true "Event"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /v1/create_event [post]
func (r *V1) create(ctx *fiber.Ctx) error {
	var body request.CreateRequest

	err := ctx.BodyParser(&body)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	if body.UserID <= 0 {
		return errorResponse(ctx, http.StatusBadRequest, "user_id required and cant be less than 1")
	}

	if body.Date == nil {
		return errorResponse(ctx, http.StatusBadRequest, "date required")
	}

	if body.Text == "" {
		return errorResponse(ctx, http.StatusBadRequest, "text required")
	}

	event := entity.Event{
		Date: body.Date.Time,
		Text: body.Text,
	}

	eventUID := uuid.New()

	err = r.e.Create(ctx.UserContext(), body.UserID, eventUID, event)
	if err != nil {
		if errors.Is(err, errs.ErrAlreadyExists) {
			return errorResponse(ctx, http.StatusInternalServerError, err.Error())
		}
		r.l.Error(err, "restapi - v1 - create")

		return errorResponse(ctx, http.StatusInternalServerError, "storage problems")
	}

	resp := response.Response{Result: response.ResultEvent{
		UID:    eventUID.String(),
		UserID: body.UserID,
		Date:   date.Date{Time: event.Date},
		Text:   event.Text,
	}}

	return ctx.Status(http.StatusOK).JSON(resp)
}

// @Summary Update
// @Description Updates event
// @ID update
// @Tags events
// @Accept json
// @Produce json
// @Param request body request.UpdateRequest true "Event"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Error
// @Failure 404 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /v1/update_event [post]
func (r *V1) update(ctx *fiber.Ctx) error {
	var body request.UpdateRequest

	err := ctx.BodyParser(&body)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	if body.UserID <= 0 {
		return errorResponse(ctx, http.StatusBadRequest, "user_id required and cant be less than 1")
	}

	if body.EventUID == "" {
		return errorResponse(ctx, http.StatusBadRequest, "uid required")
	}

	if body.Date == nil {
		return errorResponse(ctx, http.StatusBadRequest, "date required")
	}

	if body.Text == "" {
		return errorResponse(ctx, http.StatusBadRequest, "text required")
	}

	uid, err := uuid.Parse(body.EventUID)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid uid format")
	}

	err = r.e.Update(ctx.UserContext(), body.UserID, uid, body.Text, body.Date.Time)
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			return errorResponse(ctx, http.StatusNotFound, err.Error())
		} else if errors.Is(err, errs.ErrEventNotFound) {
			return errorResponse(ctx, http.StatusNotFound, err.Error())
		}
		r.l.Error(err, "restapi - v1 - update")

		return errorResponse(ctx, http.StatusInternalServerError, "storage problems")
	}

	resp := response.Response{Result: response.ResultEvent{
		UID:    body.EventUID,
		UserID: body.UserID,
		Date:   date.Date{Time: body.Date.Time},
		Text:   body.Text,
	}}

	return ctx.Status(http.StatusOK).JSON(resp)
}

// @Summary Delete
// @Description Deletes event
// @ID delete
// @Tags events
// @Accept json
// @Produce json
// @Param request body request.DeleteRequest true "Event"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Error
// @Failure 404 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /v1/delete_event [post]
func (r *V1) delete(ctx *fiber.Ctx) error {
	var body request.DeleteRequest

	err := ctx.BodyParser(&body)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid request body")
	}

	if body.UserID <= 0 {
		return errorResponse(ctx, http.StatusBadRequest, "user_id required and cant be less than 1")
	}

	if body.EventUID == "" {
		return errorResponse(ctx, http.StatusBadRequest, "uid required")
	}

	uid, err := uuid.Parse(body.EventUID)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid uid format")
	}

	err = r.e.Delete(ctx.UserContext(), body.UserID, uid)
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			return errorResponse(ctx, http.StatusNotFound, err.Error())
		} else if errors.Is(err, errs.ErrEventNotFound) {
			return errorResponse(ctx, http.StatusNotFound, err.Error())
		}
		r.l.Error(err, "restapi - v1 - delete")

		return errorResponse(ctx, http.StatusInternalServerError, "storage problems")
	}

	return ctx.SendStatus(http.StatusOK)
}

// @Summary Get events for day
// @Description Get events for day by date
// @ID get-day
// @Tags events
// @Param date query string false "Date"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Error
// @Failure 404 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /v1/events_for_day [get]
func (r *V1) getEventsForDay(ctx *fiber.Ctx) error {
	userIDStr := ctx.Query("user_id")
	dateStr := ctx.Query("date")

	d, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid date format, expected: YYYY-MM-DD")
	}

	u, err := strconv.Atoi(userIDStr)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid user_id format")
	}

	if u <= 0 {
		return errorResponse(ctx, http.StatusBadRequest, "user_id required and cant be less than 1")
	}

	events, err := r.e.GetEventsForDay(ctx.UserContext(), u, d)
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			return errorResponse(ctx, http.StatusNotFound, err.Error())
		}
		r.l.Error(err, "restapi - v1 - getEventsForDay")

		return errorResponse(ctx, http.StatusInternalServerError, "storage problems")
	}

	resps := make([]response.Response, 0, len(events))

	if len(events) == 0 {
		return ctx.Status(http.StatusOK).JSON(resps)
	}

	for uid, event := range events {
		resps = append(resps, response.Response{Result: response.ResultEvent{
			UID:    uid.String(),
			UserID: u,
			Date:   date.Date{Time: event.Date},
			Text:   event.Text,
		}})
	}

	return ctx.Status(http.StatusOK).JSON(resps)
}

// @Summary Get events for week
// @Description Get events for week by date
// @ID get-week
// @Tags events
// @Param date query string false "Date"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Error
// @Failure 404 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /v1/events_for_week [get]
func (r *V1) getEventsForWeek(ctx *fiber.Ctx) error {
	userIDStr := ctx.Query("user_id")
	dateStr := ctx.Query("date")

	d, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid date format, expected: YYYY-MM-DD")
	}

	u, err := strconv.Atoi(userIDStr)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid user_id")
	}

	if u <= 0 {
		return errorResponse(ctx, http.StatusBadRequest, "user_id required and cant be less than 1")
	}

	events, err := r.e.GetEventsForWeek(ctx.UserContext(), u, d)
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			return errorResponse(ctx, http.StatusNotFound, err.Error())
		}
		r.l.Error(err, "restapi - v1 - getEventsForWeek")

		return errorResponse(ctx, http.StatusInternalServerError, "storage problems")
	}

	resps := make([]response.Response, 0, len(events))

	if len(events) == 0 {
		return ctx.Status(http.StatusOK).JSON(resps)
	}

	for uid, event := range events {
		resps = append(resps, response.Response{Result: response.ResultEvent{
			UID:    uid.String(),
			UserID: u,
			Date:   date.Date{Time: event.Date},
			Text:   event.Text,
		}})
	}

	return ctx.Status(http.StatusOK).JSON(resps)
}

// @Summary Get events for month
// @Description Get events for month by date
// @ID get-month
// @Tags events
// @Param date query string false "Date"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Error
// @Failure 404 {object} response.Error
// @Failure 500 {object} response.Error
// @Router /v1/events_for_month [get]
func (r *V1) getEventsForMonth(ctx *fiber.Ctx) error {
	userIDStr := ctx.Query("user_id")
	dateStr := ctx.Query("date")

	d, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid date format, expected: YYYY-MM-DD")
	}

	u, err := strconv.Atoi(userIDStr)
	if err != nil {
		return errorResponse(ctx, http.StatusBadRequest, "invalid user_id")
	}

	if u <= 0 {
		return errorResponse(ctx, http.StatusBadRequest, "user_id required and cant be less than 1")
	}

	events, err := r.e.GetEventsForMonth(ctx.UserContext(), u, d)
	if err != nil {
		if errors.Is(err, errs.ErrUserNotFound) {
			return errorResponse(ctx, http.StatusNotFound, err.Error())
		}
		r.l.Error(err, "restapi - v1 - getEventsForMonth")

		return errorResponse(ctx, http.StatusInternalServerError, "storage problems")
	}

	resps := make([]response.Response, 0, len(events))

	if len(events) == 0 {
		return ctx.Status(http.StatusOK).JSON(resps)
	}

	for uid, event := range events {
		resps = append(resps, response.Response{Result: response.ResultEvent{
			UID:    uid.String(),
			UserID: u,
			Date:   date.Date{Time: event.Date},
			Text:   event.Text,
		}})
	}

	return ctx.Status(http.StatusOK).JSON(resps)
}
