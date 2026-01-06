package repo

import (
	"context"
	"time"

	"github.com/andreyxaxa/calendar/internal/entity"
	"github.com/google/uuid"
)

type (
	// EventsRepo - interface of repository
	EventsRepo interface {
		Create(ctx context.Context, userID int, eventUID uuid.UUID, event entity.Event) error
		Update(ctx context.Context, userID int, eventUID uuid.UUID, text string, date time.Time) error
		Delete(ctx context.Context, userID int, eventUID uuid.UUID) error
		GetEventsForDay(ctx context.Context, userID int, date time.Time) (map[uuid.UUID]entity.Event, error)
		GetEventsForWeek(ctx context.Context, userID int, date time.Time) (map[uuid.UUID]entity.Event, error)
		GetEventsForMonth(ctx context.Context, userID int, date time.Time) (map[uuid.UUID]entity.Event, error)
	}
)
