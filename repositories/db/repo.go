package db

import (
	"errors"

	"github.com/tuhin47/go-ems/models"
	"github.com/tuhin47/go-ems/utils/errutil"
	"gorm.io/gorm"
)

type RepositoryImpl struct {
	client *gorm.DB
}

func NewEventRepositoryImpl(client *gorm.DB) *RepositoryImpl {
	return &RepositoryImpl{
		client: client,
	}
}

func (r *RepositoryImpl) CreateEvent(event *models.Event) (*models.Event, error) {
	qry := r.client.Create(event)
	if qry.Error != nil {
		return nil, qry.Error
	}
	return event, nil
}

func (r *RepositoryImpl) ReadEventByID(id int) (*models.Event, error) {
	event := &models.Event{}
	qry := r.client.First(event, "id = ?", id)
	if errors.Is(qry.Error, gorm.ErrRecordNotFound) {
		return nil, errutil.ErrRecordNotFound
	}
	if qry.Error != nil {
		return nil, qry.Error
	}
	return event, nil
}

func (r *RepositoryImpl) ListEvents() ([]*models.Event, error) {
	var events []*models.Event
	qry := r.client.Find(&events)
	if qry.RowsAffected == 0 {
		return nil, errutil.ErrRecordNotFound
	}
	if qry.Error != nil {
		return nil, qry.Error
	}
	return events, nil
}
