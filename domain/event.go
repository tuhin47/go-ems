package domain

import (
	"github.com/tuhin47/go-ems/models"
	"github.com/tuhin47/go-ems/types"
)

type EventRepository interface {
	CreateEvent(event *models.Event) (*models.Event, error)
	ReadEventByID(id int) (*models.Event, error)
	ListEvents() ([]*models.Event, error)
}

type EventService interface {
	CreateEvent(req *types.EventCreateRequest) (*types.EventCreateResponse, error)
	ReadEventByID(id int) (*models.Event, error)
	ListEvents() ([]*models.Event, error)
}
