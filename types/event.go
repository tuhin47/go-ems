package types

import (
	"time"

	"github.com/tuhin47/go-ems/models"
)

type EventReadRequest struct {
	ID int `param:"id"`
}

type EventUpsertRequest struct {
	ID          *int    `param:"id"`
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

func (req *EventUpsertRequest) ToEvent() *models.Event {
	event := &models.Event{
		Title:       req.Title,
		Description: req.Description,
		Location:    req.Location,
		CreatedBy:   req.CreatedBy,
	}
	if req.StartTime != nil && len(*req.StartTime) > 0 {
		t, _ := time.Parse(time.RFC3339, *req.StartTime)
		event.StartTime = &t
	}
	if req.EndTime != nil && len(*req.EndTime) > 0 {
		t, _ := time.Parse(time.RFC3339, *req.EndTime)
		event.EndTime = &t
	}
	return event
}
