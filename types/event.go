package types

import (
	"time"

	"github.com/vivasoft-ltd/go-ems/models"
)

type EventReadRequest struct {
	ID int `param:"id"`
}

type EventCreateRequest struct {
	Title       string  `json:"title"`
	Description *string `json:"description"`
	Location    *string `json:"location"`
	StartTime   *string `json:"start_time"`
	EndTime     *string `json:"end_time"`
	CreatedBy   *string `json:"created_by"`
}

type EventCreateResponse struct {
	Message string        `json:"message"`
	Event   *models.Event `json:"event"`
}

func (ecreq *EventCreateRequest) ToEvent() *models.Event {
	event := &models.Event{
		Title:       ecreq.Title,
		Description: ecreq.Description,
		Location:    ecreq.Location,
		CreatedBy:   ecreq.CreatedBy,
	}

	if ecreq.StartTime != nil && len(*ecreq.StartTime) > 0 {
		t, _ := time.Parse(time.RFC3339, *ecreq.StartTime)
		event.StartTime = &t
	}
	if ecreq.EndTime != nil && len(*ecreq.EndTime) > 0 {
		t, _ := time.Parse(time.RFC3339, *ecreq.EndTime)
		event.EndTime = &t
	}
	return event
}
