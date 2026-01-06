package v1

import (
	"github.com/andreyxaxa/calendar/internal/usecase"
	"github.com/andreyxaxa/calendar/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

// NewEventsRoutes -.
func NewEventsRoutes(apiV1Group fiber.Router, e usecase.Events, l logger.Interface) {
	r := &V1{
		e: e,
		l: l,
	}

	{
		apiV1Group.Post("/create_event", r.create)
		apiV1Group.Post("/update_event", r.update)
		apiV1Group.Post("/delete_event", r.delete)

		apiV1Group.Get("/events_for_day", r.getEventsForDay)
		apiV1Group.Get("/events_for_week", r.getEventsForWeek)
		apiV1Group.Get("/events_for_month", r.getEventsForMonth)
	}
}
