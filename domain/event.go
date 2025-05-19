package domain

import (
	"github.com/vivasoft-ltd/go-ems/models"
	"github.com/vivasoft-ltd/go-ems/types"
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
