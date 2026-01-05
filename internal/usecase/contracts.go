package usecase

import (
	"context"
	"time"

	"github.com/andreyxaxa/calendar/internal/entity"
	"github.com/google/uuid"
)

type (
	Events interface {
		Create(context.Context, int, uuid.UUID, entity.Event) error
		Update(context.Context, int, uuid.UUID, string, time.Time) error
		Delete(context.Context, int, uuid.UUID) error
		GetEventsForDay(context.Context, int, time.Time) (map[uuid.UUID]entity.Event, error)
		GetEventsForWeek(context.Context, int, time.Time) (map[uuid.UUID]entity.Event, error)
		GetEventsForMonth(context.Context, int, time.Time) (map[uuid.UUID]entity.Event, error)
	}
)
