package controllers

import (
	"github.com/labstack/echo"
	"net/http"
	"service-template-golang/domain/entity"
	"service-template-golang/http/services"
	env "service-template-golang/utils"
)

// EventsController encapsulates all the data necessary for the implementation of the EventsService.
type EventsController struct {
	eventsService services.EventsService
}

// NewEventsController instantiate a new event controller.
func NewEventsController(es services.EventsService) *EventsController {
	return &EventsController{
		eventsService: es,
	}
}

// GetID return a event by ID [eventsService.GetID].
func (ec *EventsController) GetID(c echo.Context) error {
	ID, err := env.GetParam(c, "id")
	if err != nil {
		return entity.NewError(http.StatusBadRequest, err)
	}
	events, err := ec.eventsService.GetID(ID)
	if err != nil {
		return entity.HandleServiceError(err)
	}
	return c.JSON(http.StatusOK, events.ToDomainEvents())
}
