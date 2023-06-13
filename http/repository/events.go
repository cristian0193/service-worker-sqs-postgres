package repository

import (
	"service-template-golang/database"
	"service-template-golang/domain/entity"
)

// EventsRepository encapsulates all the data needed to the persistence in the event table.
type EventsRepository struct {
	db *database.ClientDB
}

// NewEventsRepository instance the connection to the database and use of log.
func NewEventsRepository(db *database.ClientDB) *EventsRepository {
	return &EventsRepository{
		db: db,
	}
}

// GetID return the last store by ID.
func (er *EventsRepository) GetID(ID string) (*entity.Events, error) {
	event := &entity.Events{}

	err := er.db.DB.Model(&event).Where("id = ?", ID).Scan(&event).Error
	if err != nil {
		return nil, entity.ErrInternalError
	}

	return event, nil
}
