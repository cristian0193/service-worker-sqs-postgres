package events

import (
	"service-worker-sqs-postgres/core/domain/entity"
	"service-worker-sqs-postgres/core/domain/exceptions"
	"service-worker-sqs-postgres/dataproviders/database"
)

type IEventsRepository interface {
	GetID(ID string) (*entity.Events, error)
}

// EventsRepository encapsulates all the data needed to the persistence in the event table.
type EventsRepository struct {
	db *database.ClientDB
}

// NewEventsRepository instance the connection to the database.
func NewEventsRepository(db *database.ClientDB) *EventsRepository {
	return &EventsRepository{
		db: db,
	}
}

// GetID return the event by ID.
func (er *EventsRepository) GetID(ID string) (*entity.Events, error) {
	event := &entity.Events{}

	err := er.db.DB.Model(&event).Where("id = ?", ID).Scan(&event).Error
	if err != nil {
		return nil, exceptions.ErrInternalError
	}

	return event, nil
}
