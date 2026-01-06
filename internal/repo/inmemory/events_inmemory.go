package inmemory

import (
	"context"
	"sync"
	"time"

	"github.com/andreyxaxa/calendar/internal/entity"
	"github.com/andreyxaxa/calendar/pkg/types/errs"
	"github.com/google/uuid"
)

// EventsRepo -.
type EventsRepo struct {
	storage map[int]map[uuid.UUID]entity.Event
	mu      sync.RWMutex
}

// New returns new EventsRepo(struct)
func New() *EventsRepo {
	return &EventsRepo{
		storage: make(map[int]map[uuid.UUID]entity.Event),
	}
}

// Create -.
func (r *EventsRepo) Create(ctx context.Context, userID int, eventUID uuid.UUID, event entity.Event) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.storage[userID]; !ok {
		r.storage[userID] = make(map[uuid.UUID]entity.Event)
	}

	if _, ok := r.storage[userID][eventUID]; ok {
		return errs.ErrAlreadyExists
	}

	r.storage[userID][eventUID] = event
	return nil
}

// Update -.
func (r *EventsRepo) Update(ctx context.Context, userID int, eventUID uuid.UUID, text string, date time.Time) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.storage[userID]; !ok {
		return errs.ErrUserNotFound
	}

	event, ok := r.storage[userID][eventUID]
	if !ok {
		return errs.ErrEventNotFound
	}

	event.Text = text
	event.Date = date
	r.storage[userID][eventUID] = event

	return nil
}

// Delete -.
func (r *EventsRepo) Delete(ctx context.Context, userID int, eventUID uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, ok := r.storage[userID]; !ok {
		return errs.ErrUserNotFound
	}

	_, ok := r.storage[userID][eventUID]
	if !ok {
		return errs.ErrEventNotFound
	}

	delete(r.storage[userID], eventUID)

	return nil
}

// GetEventsForDay -.
func (r *EventsRepo) GetEventsForDay(ctx context.Context, userID int, date time.Time) (map[uuid.UUID]entity.Event, error) {
	events := make(map[uuid.UUID]entity.Event)

	r.mu.RLock()
	userEvents, ok := r.storage[userID]
	if !ok {
		return nil, errs.ErrUserNotFound
	}

	for uid, event := range userEvents {
		if event.Date.Equal(date) {
			events[uid] = event
		}
	}
	r.mu.RUnlock()

	return events, nil
}

// GetEventsForWeek -.
func (r *EventsRepo) GetEventsForWeek(ctx context.Context, userID int, date time.Time) (map[uuid.UUID]entity.Event, error) {
	events := make(map[uuid.UUID]entity.Event)

	r.mu.RLock()
	userEvents, ok := r.storage[userID]
	if !ok {
		return nil, errs.ErrUserNotFound
	}

	targetYear, targetWeek := date.ISOWeek()

	for uid, event := range userEvents {
		actualYear, actualWeek := event.Date.ISOWeek()

		if actualYear == targetYear && actualWeek == targetWeek {
			events[uid] = event
		}
	}
	r.mu.RUnlock()

	return events, nil
}

// GetEventsForMonth -.
func (r *EventsRepo) GetEventsForMonth(ctx context.Context, userID int, date time.Time) (map[uuid.UUID]entity.Event, error) {
	events := make(map[uuid.UUID]entity.Event)

	r.mu.RLock()

	userEvents, ok := r.storage[userID]
	if !ok {
		return nil, errs.ErrUserNotFound
	}

	targetYear, targetMonth, _ := date.Date()

	for uid, event := range userEvents {
		actualYear, actualMonth, _ := event.Date.Date()
		if actualYear == targetYear && actualMonth == targetMonth {
			events[uid] = event
		}
	}
	r.mu.RUnlock()

	return events, nil
}
