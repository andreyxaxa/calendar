package v1

import (
	"github.com/andreyxaxa/calendar/internal/usecase"
	"github.com/andreyxaxa/calendar/pkg/logger"
)

type V1 struct {
	l logger.Interface
	e usecase.Events
}
