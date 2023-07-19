package events

import (
	"gorm.io/gorm/clause"
	"service-worker-sqs-postgres/core/domain/entity"
	"service-worker-sqs-postgres/core/domain/exceptions"
	"service-worker-sqs-postgres/dataproviders/postgres"
)

type IEventsRepository interface {
	GetID(ID string) (*entity.Events, error)
	Insert(events *entity.Events) error
}

// EventsRepository encapsulates all the data needed to the persistence in the event table.
type EventsRepository struct {
	db *postgres.ClientDB
}

// NewEventsRepository instance the connection to the postgres.
func NewEventsRepository(db *postgres.ClientDB) *EventsRepository {
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

// Insert records an event in the database.
func (er *EventsRepository) Insert(events *entity.Events) error {
	r := er.db.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&events)
	if r.Error != nil {
		r.Rollback()
		return r.Error
	}
	return nil
}
