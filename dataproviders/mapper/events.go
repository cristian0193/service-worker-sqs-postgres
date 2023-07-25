package mapper

import (
	"service-worker-sqs-postgres/core/domain"
	"service-worker-sqs-postgres/core/domain/entity"
)

// ToDomainEvents convert domain event to model the postgres events .
func ToDomainEvents(e *entity.Events) *domain.Events {
	return &domain.Events{
		ID:      e.ID,
		Message: e.Message,
		Date:    e.Date,
	}
}

// ToEntityEvents convert entity event to model the postgres events .
func ToEntityEvents(e *domain.Events) *entity.Events {
	return &entity.Events{
		ID:      e.ID,
		Message: e.Message,
		Date:    e.Date,
	}
}
