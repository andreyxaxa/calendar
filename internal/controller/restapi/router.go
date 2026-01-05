package restapi

import (
	"github.com/andreyxaxa/calendar/config"
	"github.com/andreyxaxa/calendar/internal/controller/restapi/middleware"
	v1 "github.com/andreyxaxa/calendar/internal/controller/restapi/v1"
	"github.com/andreyxaxa/calendar/internal/usecase"
	"github.com/andreyxaxa/calendar/pkg/logger"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(app *fiber.App, cfg *config.Config, e usecase.Events, l logger.Interface) {
	// Options
	app.Use(middleware.Logger(l))

	apiV1Group := app.Group("/v1")
	{
		v1.NewEventsRoutes(apiV1Group, e, l)
	}
}
