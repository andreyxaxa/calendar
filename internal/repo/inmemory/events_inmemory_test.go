package inmemory_test

import (
	"context"
	"testing"
	"time"

	"github.com/andreyxaxa/calendar/internal/entity"
	"github.com/andreyxaxa/calendar/internal/repo/inmemory"
	"github.com/google/uuid"
)

func TestCreateAndGetForDay(t *testing.T) {
	repo := inmemory.New()

	ctx := context.Background()
	userID := 1
	date := time.Date(2026, 1, 1, 10, 0, 0, 0, time.UTC)
	uid := uuid.New()

	event := entity.Event{
		Text: "meeting with friend",
		Date: date,
	}

	err := repo.Create(ctx, userID, uid, event)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	events, err := repo.GetEventsForDay(ctx, userID, date)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(events) != 1 {
		t.Fatalf("expected 1 event, got %d", len(events))
	}

	expected := event.Text
	actual := events[uid].Text

	if actual != expected {
		t.Fatalf("expected %q, got %q", expected, actual)
	}
}

func TestCreateAndUpdate(t *testing.T) {
	repo := inmemory.New()

	ctx := context.Background()
	userID := 1
	uid := uuid.New()
	date := time.Now()

	newDate := date.Add(24 * time.Hour)

	err := repo.Create(ctx, userID, uid, entity.Event{
		Text: "old",
		Date: date,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = repo.Update(ctx, userID, uid, "new", newDate)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// запрашиваю со старой датой
	events, err := repo.GetEventsForDay(ctx, userID, date)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// ожидаю пустой ответ
	if len(events) != 0 {
		t.Fatalf("expected 0 events, got %d", len(events))
	}

	// запрашиваю с новой датой
	events, err = repo.GetEventsForDay(ctx, userID, newDate)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if events[uid].Text != "new" {
		t.Fatalf("expected updated text, got old")
	}
}

func TestCreateAndDelete(t *testing.T) {
	repo := inmemory.New()

	ctx := context.Background()
	userID := 1
	uid := uuid.New()
	date := time.Now()

	err := repo.Create(ctx, userID, uid, entity.Event{
		Text: "meeting with friend",
		Date: date,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	err = repo.Delete(ctx, userID, uid)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	events, err := repo.GetEventsForDay(ctx, userID, date)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(events) != 0 {
		t.Fatalf("expected 0 events, got %d", len(events))
	}
}
