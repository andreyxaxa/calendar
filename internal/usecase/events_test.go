package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/andreyxaxa/calendar/internal/entity"
	"github.com/andreyxaxa/calendar/internal/usecase/events"
	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
)

var errStorageProblem = errors.New("storage problems")

func eventsUseCase(t *testing.T) (*events.UseCase, *MockEventsRepo, *gomock.Controller) {
	t.Helper()

	mockCtl := gomock.NewController(t)

	repo := NewMockEventsRepo(mockCtl)

	useCase := events.New(repo)

	return useCase, repo, mockCtl
}

func TestCreateOK(t *testing.T) {
	t.Parallel()

	useCase, repo, ctrl := eventsUseCase(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := 1
	eventUID := uuid.New()
	event := entity.Event{
		Text: "meeting with friend",
		Date: time.Now(),
	}

	repo.
		EXPECT().
		Create(ctx, userID, eventUID, event).
		Return(nil)

	err := useCase.Create(ctx, userID, eventUID, event)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestCreateErr(t *testing.T) {
	t.Parallel()

	useCase, repo, ctrl := eventsUseCase(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := 1
	eventUID := uuid.New()
	event := entity.Event{}

	repo.
		EXPECT().
		Create(ctx, userID, eventUID, event).
		Return(errStorageProblem)

	err := useCase.Create(ctx, userID, eventUID, event)

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, errStorageProblem) {
		t.Fatalf("expected wrapped error, got %v", err)
	}
}

func TestUpdateOK(t *testing.T) {
	t.Parallel()

	useCase, repo, ctrl := eventsUseCase(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := 1
	eventUID := uuid.New()
	text := "updated text"
	date := time.Now()

	repo.
		EXPECT().
		Update(ctx, userID, eventUID, text, date).
		Return(nil)

	err := useCase.Update(ctx, userID, eventUID, text, date)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUpdateErr(t *testing.T) {
	t.Parallel()

	useCase, repo, ctrl := eventsUseCase(t)
	defer ctrl.Finish()

	repo.
		EXPECT().
		Update(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		Return(errStorageProblem)

	err := useCase.Update(context.Background(), 1, uuid.New(), "text", time.Now())

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, errStorageProblem) {
		t.Fatalf("expected wrapped error, got %v", err)
	}
}

func TestDeleteOK(t *testing.T) {
	t.Parallel()

	useCase, repo, ctrl := eventsUseCase(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := 1
	eventUID := uuid.New()

	repo.
		EXPECT().
		Delete(ctx, userID, eventUID).
		Return(nil)

	err := useCase.Delete(ctx, userID, eventUID)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestDeleteErr(t *testing.T) {
	t.Parallel()

	useCase, repo, ctrl := eventsUseCase(t)
	defer ctrl.Finish()

	repo.
		EXPECT().
		Delete(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(errStorageProblem)

	err := useCase.Delete(context.Background(), 1, uuid.New())

	if err == nil {
		t.Fatal("expected error")
	}

	if !errors.Is(err, errStorageProblem) {
		t.Fatalf("expected wrapped error, got %v", err)
	}
}

func TestGetEventsForDayOK(t *testing.T) {
	t.Parallel()

	useCase, repo, ctrl := eventsUseCase(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := 1
	date := time.Now()

	expected := map[uuid.UUID]entity.Event{uuid.New(): {Text: "text"}}

	repo.
		EXPECT().
		GetEventsForDay(ctx, userID, date).
		Return(expected, nil)

	result, err := useCase.GetEventsForDay(ctx, userID, date)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != len(expected) {
		t.Fatalf("expected %d events, got %d", len(expected), len(result))
	}
}

func TestGetEventsForDayErr(t *testing.T) {
	t.Parallel()

	useCase, repo, ctrl := eventsUseCase(t)
	defer ctrl.Finish()

	repo.
		EXPECT().
		GetEventsForDay(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, errStorageProblem)

	result, err := useCase.GetEventsForDay(context.Background(), 1, time.Now())

	if err == nil {
		t.Fatal("expected error")
	}

	if result != nil {
		t.Fatal("expected nil result")
	}

	if !errors.Is(err, errStorageProblem) {
		t.Fatalf("expected wrapped error, got %v", err)
	}
}

func TestGetEventsForWeekOK(t *testing.T) {
	t.Parallel()

	useCase, repo, ctrl := eventsUseCase(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := 1
	date := time.Now()

	expected := map[uuid.UUID]entity.Event{uuid.New(): {Text: "text"}}

	repo.
		EXPECT().
		GetEventsForWeek(ctx, userID, date).
		Return(expected, nil)

	result, err := useCase.GetEventsForWeek(ctx, userID, date)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != len(expected) {
		t.Fatalf("expected %d events, got %d", len(expected), len(result))
	}
}

func TestGetEventsForWeekErr(t *testing.T) {
	t.Parallel()

	useCase, repo, ctrl := eventsUseCase(t)
	defer ctrl.Finish()

	repo.
		EXPECT().
		GetEventsForWeek(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, errStorageProblem)

	result, err := useCase.GetEventsForWeek(context.Background(), 1, time.Now())

	if err == nil {
		t.Fatal("expected error")
	}

	if result != nil {
		t.Fatal("expected nil result")
	}

	if !errors.Is(err, errStorageProblem) {
		t.Fatalf("expected wrapped error, got %v", err)
	}
}

func TestGetEventsForMonthOK(t *testing.T) {
	t.Parallel()

	useCase, repo, ctrl := eventsUseCase(t)
	defer ctrl.Finish()

	ctx := context.Background()
	userID := 1
	date := time.Now()

	expected := map[uuid.UUID]entity.Event{uuid.New(): {Text: "text"}}

	repo.
		EXPECT().
		GetEventsForMonth(ctx, userID, date).
		Return(expected, nil)

	result, err := useCase.GetEventsForMonth(ctx, userID, date)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(result) != len(expected) {
		t.Fatalf("expected %d events, got %d", len(expected), len(result))
	}
}

func TestGetEventsForMonthErr(t *testing.T) {
	t.Parallel()

	useCase, repo, ctrl := eventsUseCase(t)
	defer ctrl.Finish()

	repo.
		EXPECT().
		GetEventsForMonth(gomock.Any(), gomock.Any(), gomock.Any()).
		Return(nil, errStorageProblem)

	result, err := useCase.GetEventsForMonth(context.Background(), 1, time.Now())

	if err == nil {
		t.Fatal("expected error")
	}

	if result != nil {
		t.Fatal("expected nil result")
	}

	if !errors.Is(err, errStorageProblem) {
		t.Fatalf("expected wrapped error, got %v", err)
	}
}
