package services

import (
	"errors"
	"fmt"

	"github.com/vivasoft-ltd/go-ems/domain"
	"github.com/vivasoft-ltd/go-ems/models"
	"github.com/vivasoft-ltd/go-ems/types"
	"github.com/vivasoft-ltd/go-ems/utils/errutil"
	"github.com/vivasoft-ltd/golang-course-utils/logger"
)

type EventServiceImpl struct {
	eventRepo domain.EventRepository
}

func NewEventServiceImpl(eventRepo domain.EventRepository) *EventServiceImpl {
	return &EventServiceImpl{
		eventRepo: eventRepo,
	}
}

func (s *EventServiceImpl) CreateEvent(req *types.EventCreateRequest) (*types.EventCreateResponse, error) {
	event := req.ToEvent()
	createdEvent, err := s.eventRepo.CreateEvent(event)
	if err != nil {
		logger.Error(fmt.Sprintf("err: [%v] occurred while creating event", err))
		return nil, err
	}
	return &types.EventCreateResponse{
		Message: "Event created successfully",
		Event:   createdEvent,
	}, nil
}

func (s *EventServiceImpl) ReadEventByID(id int) (*models.Event, error) {
	event, err := s.eventRepo.ReadEventByID(id)
	if err != nil {
		logger.Error(fmt.Sprintf("err: [%v] occurred while reading event by id", err))
		return nil, err
	}
	return event, nil
}

func (s *EventServiceImpl) ListEvents() ([]*models.Event, error) {
	events, err := s.eventRepo.ListEvents()
	if errors.Is(err, errutil.ErrRecordNotFound) {
		logger.Error(fmt.Sprintf("err: [%v] occurred while listing events", err))
		return []*models.Event{}, nil
	}
	if err != nil {
		logger.Error(fmt.Sprintf("err: [%v] occurred while listing events", err))
		return nil, err
	}
	return events, nil
}
