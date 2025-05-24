package services

import (
	"errors"
	"github.com/tuhin47/go-ems/domain"
	"github.com/tuhin47/go-ems/models"
	"github.com/tuhin47/go-ems/types"
	"github.com/tuhin47/go-ems/utils/errutil"
	//"github.com/tuhin47/golang-course-utils/logger"
	"time"
)

type EventServiceImpl struct {
	eventRepo domain.EventRepository
}

func NewEventServiceImpl(eventRepo domain.EventRepository) *EventServiceImpl {
	return &EventServiceImpl{
		eventRepo: eventRepo,
	}
}

func (s *EventServiceImpl) CreateEvent(req *types.EventUpsertRequest) (*types.EventCreateResponse, error) {
	event := req.ToEvent()
	createdEvent, err := s.eventRepo.CreateEvent(event)
	if err != nil {
		//logger.Error(fmt.Sprintf("err: [%v] occurred while creating event", err))
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
		//logger.Error(fmt.Sprintf("err: [%v] occurred while reading event by id", err))
		return nil, err
	}
	return event, nil
}

func (s *EventServiceImpl) ListEvents() ([]*models.Event, error) {
	events, err := s.eventRepo.ListEvents()
	if errors.Is(err, errutil.ErrRecordNotFound) {
		//logger.Error(fmt.Sprintf("err: [%v] occurred while listing events", err))
		return []*models.Event{}, nil
	}
	if err != nil {
		//logger.Error(fmt.Sprintf("err: [%v] occurred while listing events", err))
		return nil, err
	}
	return events, nil
}

func (s *EventServiceImpl) DeleteEvent(id int) error {
	err := s.eventRepo.DeleteEvent(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *EventServiceImpl) UpdateEvent(req *types.EventUpsertRequest) (*models.Event, error) {
	if req.ID == nil {
		return nil, errors.New("missing event ID for update")
	}
	event, err := s.eventRepo.ReadEventByID(*req.ID)
	if err != nil {
		return nil, err
	}
	// Update fields
	event.Title = req.Title
	event.Description = req.Description
	event.Location = req.Location
	if req.StartTime != nil && len(*req.StartTime) > 0 {
		t, _ := time.Parse(time.RFC3339, *req.StartTime)
		event.StartTime = &t
	}
	if req.EndTime != nil && len(*req.EndTime) > 0 {
		t, _ := time.Parse(time.RFC3339, *req.EndTime)
		event.EndTime = &t
	}
	event.CreatedBy = req.CreatedBy

	updatedEvent, err := s.eventRepo.UpdateEvent(event)
	if err != nil {
		return nil, err
	}
	return updatedEvent, nil
}
