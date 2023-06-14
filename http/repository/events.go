package repository

import "service-worker-sqs-postgres/domain/entity"

// EventsRepository implementation of interfaces.
type EventsRepository interface {
	GetID(ID string) (*entity.Events, error)
}
