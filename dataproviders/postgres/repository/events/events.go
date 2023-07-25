package repository

import (
	"gorm.io/gorm/clause"
	"service-worker-sqs-postgres/core/domain"
	"service-worker-sqs-postgres/core/domain/entity"
	"service-worker-sqs-postgres/core/domain/exceptions"
	"service-worker-sqs-postgres/dataproviders/mapper"
	"service-worker-sqs-postgres/dataproviders/postgres"
)

// IEventRepository interface by repository.
type IEventRepository interface {
	GetID(ID string) (*domain.Events, error)
	Insert(events *domain.Events) error
}

// EventRepository encapsulates all the data needed to the persistence in the event table.
type EventRepository struct {
	db *postgres.ClientDB
}

// NewEventRepository instance the connection to the postgres.
func NewEventRepository(db *postgres.ClientDB) *EventRepository {
	return &EventRepository{
		db: db,
	}
}

// GetID return the event by ID.
func (er *EventRepository) GetID(ID string) (*domain.Events, error) {
	event := &entity.Events{}

	err := er.db.DB.Model(&event).Where("id = ?", ID).Scan(&event).Error
	if err != nil {
		return nil, exceptions.ErrInternalError
	}

	return mapper.ToDomainEvents(event), nil
}

// Insert records an event in the database.
func (er *EventRepository) Insert(events *domain.Events) error {

	event := mapper.ToEntityEvents(events)

	r := er.db.DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&event)
	if r.Error != nil {
		r.Rollback()
		return r.Error
	}
	return nil
}
