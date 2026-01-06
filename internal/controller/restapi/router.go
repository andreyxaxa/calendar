package restapi

import (
	"github.com/andreyxaxa/calendar/config"
	_ "github.com/andreyxaxa/calendar/docs" // Swagger docs.
	"github.com/andreyxaxa/calendar/internal/controller/restapi/middleware"
	v1 "github.com/andreyxaxa/calendar/internal/controller/restapi/v1"
	"github.com/andreyxaxa/calendar/internal/usecase"
	"github.com/andreyxaxa/calendar/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// NewRouter -.
// @title HTTP-Calendar
// @version 1.0
// @host localhost:8080
// @BasePath /v1
func NewRouter(app *fiber.App, cfg *config.Config, e usecase.Events, l logger.Interface) {
	// Swagger
	if cfg.Swagger.Enabled {
		app.Get("/swagger/*", swagger.HandlerDefault)
	}

	// Options
	app.Use(middleware.Logger(l))

	// Routers
	apiV1Group := app.Group("/v1")
	{
		v1.NewEventsRoutes(apiV1Group, e, l)
	}
}
