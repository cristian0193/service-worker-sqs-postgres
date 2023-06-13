package repository

import "service-template-golang/domain/entity"

// EventsRepository implementation of interfaces.
type EventsRepository interface {
	GetID(ID string) (*entity.Events, error)
}
