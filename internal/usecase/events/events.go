package events

import (
	"context"
	"fmt"
	"time"

	"github.com/andreyxaxa/calendar/internal/entity"
	"github.com/andreyxaxa/calendar/internal/repo"
	"github.com/google/uuid"
)

// UseCase -.
type UseCase struct {
	repo repo.EventsRepo
}

// New returns new UseCase(struct)
func New(r repo.EventsRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

// Create -.
func (uc *UseCase) Create(ctx context.Context, userID int, eventUID uuid.UUID, event entity.Event) error {
	if err := uc.repo.Create(ctx, userID, eventUID, event); err != nil {
		return fmt.Errorf("EventsUseCase - Create - uc.repo.Create: %w", err)
	}

	return nil
}

// Update -.
func (uc *UseCase) Update(ctx context.Context, userID int, eventUID uuid.UUID, text string, date time.Time) error {
	if err := uc.repo.Update(ctx, userID, eventUID, text, date); err != nil {
		return fmt.Errorf("EventsUseCase - Update - uc.repo.Update: %w", err)
	}

	return nil
}

// Delete -.
func (uc *UseCase) Delete(ctx context.Context, userID int, eventUID uuid.UUID) error {
	if err := uc.repo.Delete(ctx, userID, eventUID); err != nil {
		return fmt.Errorf("EventsUseCase - Delete - uc.repo.Delete: %w", err)
	}

	return nil
}

// GetEventsForDay -.
func (uc *UseCase) GetEventsForDay(ctx context.Context, userID int, date time.Time) (map[uuid.UUID]entity.Event, error) {
	events, err := uc.repo.GetEventsForDay(ctx, userID, date)
	if err != nil {
		return nil, fmt.Errorf("EventsUseCase - GetEventsForDay - uc.repo.GetEventsForDay: %w", err)
	}

	return events, nil
}

// GetEventsForWeek -.
func (uc *UseCase) GetEventsForWeek(ctx context.Context, userID int, date time.Time) (map[uuid.UUID]entity.Event, error) {
	events, err := uc.repo.GetEventsForWeek(ctx, userID, date)
	if err != nil {
		return nil, fmt.Errorf("EventsUseCase - GetEventsForWeek - uc.repo.GetEventsForWeek: %w", err)
	}

	return events, nil
}

// GetEventsForMonth -.
func (uc *UseCase) GetEventsForMonth(ctx context.Context, userID int, date time.Time) (map[uuid.UUID]entity.Event, error) {
	events, err := uc.repo.GetEventsForMonth(ctx, userID, date)
	if err != nil {
		return nil, fmt.Errorf("EventsUseCase - GetEventsForMonth - uc.repo.GetEventsForMonth: %w", err)
	}

	return events, nil
}
