package domain

import (
	"github.com/tuhin47/go-ems/models"
	"github.com/tuhin47/go-ems/types"
)

type EventRepository interface {
	CreateEvent(event *models.Event) (*models.Event, error)
	ReadEventByID(id int) (*models.Event, error)
	ListEvents() ([]*models.Event, error)
	DeleteEvent(id int) error
	UpdateEvent(event *models.Event) (*models.Event, error)
}

type EventService interface {
	CreateEvent(req *types.EventUpsertRequest) (*types.EventCreateResponse, error)
	ReadEventByID(id int) (*models.Event, error)
	ListEvents() ([]*models.Event, error)
	DeleteEvent(id int) error
	UpdateEvent(req *types.EventUpsertRequest) (*models.Event, error)
}
