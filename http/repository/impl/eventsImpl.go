package impl

import (
	"service-template-golang/database"
	"service-template-golang/domain/entity"
)

// EventsRepositoryImpl encapsulates all the data needed to the persistence in the event table.
type EventsRepositoryImpl struct {
	db *database.ClientDB
}

// NewEventsRepository instance the connection to the database.
func NewEventsRepository(db *database.ClientDB) *EventsRepositoryImpl {
	return &EventsRepositoryImpl{
		db: db,
	}
}

// GetID return the event by ID.
func (er *EventsRepositoryImpl) GetID(ID string) (*entity.Events, error) {
	event := &entity.Events{}

	err := er.db.DB.Model(&event).Where("id = ?", ID).Scan(&event).Error
	if err != nil {
		return nil, entity.ErrInternalError
	}

	return event, nil
}
