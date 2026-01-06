package v1

import (
	"github.com/andreyxaxa/calendar/internal/usecase"
	"github.com/andreyxaxa/calendar/pkg/logger"
)

// V1 -.
type V1 struct {
	l logger.Interface
	e usecase.Events
}
