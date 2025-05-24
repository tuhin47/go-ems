package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/tuhin47/go-ems/domain"
	"github.com/tuhin47/go-ems/types"
	"github.com/tuhin47/go-ems/utils/errutil"
	"github.com/tuhin47/go-ems/utils/msgutil"
)

type EventController struct {
	eventSvc domain.EventService
}

func NewEventController(eventSvc domain.EventService) *EventController {
	return &EventController{
		eventSvc: eventSvc,
	}
}

func (ec *EventController) CreateEvent(c echo.Context) error {
	var req types.EventUpsertRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, msgutil.RequestBodyParseError())
	}

	resp, err := ec.eventSvc.CreateEvent(&req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msgutil.SomethingWentWrong())
	}
	return c.JSON(http.StatusCreated, resp)
}

func (ec *EventController) ReadEventByID(e echo.Context) error {
	var eventReadReq types.EventReadRequest
	if err := e.Bind(&eventReadReq); err != nil {
		return e.JSON(http.StatusBadRequest, msgutil.RequestBodyParseError())
	}

	event, err := ec.eventSvc.ReadEventByID(eventReadReq.ID)
	if errors.Is(err, errutil.ErrRecordNotFound) {
		return e.JSON(http.StatusNotFound, msgutil.RecordNotFound())
	}
	if err != nil {
		return e.JSON(http.StatusInternalServerError, msgutil.SomethingWentWrong())
	}
	return e.JSON(http.StatusOK, event)
}

func (ec *EventController) ListEvents(c echo.Context) error {
	events, err := ec.eventSvc.ListEvents()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, msgutil.SomethingWentWrong())
	}
	return c.JSON(http.StatusOK, events)
}

func (ec *EventController) DeleteEvent(e echo.Context) error {
	var eventReadReq types.EventReadRequest
	if err := e.Bind(&eventReadReq); err != nil {
		return e.JSON(http.StatusBadRequest, msgutil.RequestBodyParseError())
	}
	err := ec.eventSvc.DeleteEvent(eventReadReq.ID)
	if errors.Is(err, errutil.ErrRecordNotFound) {
		return e.JSON(http.StatusNotFound, msgutil.RecordNotFound())
	}
	if err != nil {
		return e.JSON(http.StatusInternalServerError, msgutil.SomethingWentWrong())
	}
	return e.JSON(http.StatusOK, msgutil.EventDeletedSuccessfully())
}

func (ec *EventController) UpdateEvent(e echo.Context) error {
	var req types.EventUpsertRequest
	if err := e.Bind(&req); err != nil {
		return e.JSON(http.StatusBadRequest, msgutil.RequestBodyParseError())
	}
	updatedEvent, err := ec.eventSvc.UpdateEvent(&req)
	if errors.Is(err, errutil.ErrRecordNotFound) {
		return e.JSON(http.StatusNotFound, msgutil.RecordNotFound())
	}
	if err != nil {
		return e.JSON(http.StatusInternalServerError, msgutil.SomethingWentWrong())
	}
	return e.JSON(http.StatusOK, updatedEvent)
}
