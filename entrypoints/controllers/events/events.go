package events

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"service-worker-sqs-postgres/core/domain/exceptions"
	cases "service-worker-sqs-postgres/core/usecases/events"
	env "service-worker-sqs-postgres/dataproviders/utils"
)

// EventsController encapsulates all the data necessary for the implementation of the EventsService.
type EventsController struct {
	eventsUseCases cases.IEventsCaseUses
}

// NewEventsController instantiate a new event controller.
func NewEventsController(es cases.IEventsCaseUses) *EventsController {
	return &EventsController{
		eventsUseCases: es,
	}
}

// GetID return a event by ID [eventsService.GetID].
func (ec *EventsController) GetID(c echo.Context) error {
	ID, err := env.GetParam(c, "id")
	if err != nil {
		return exceptions.NewError(http.StatusBadRequest, err)
	}
	events, err := ec.eventsUseCases.GetID(ID)
	if err != nil {
		return exceptions.HandleServiceError(err)
	}
	return c.JSON(http.StatusOK, events.ToDomainEvents())
}
