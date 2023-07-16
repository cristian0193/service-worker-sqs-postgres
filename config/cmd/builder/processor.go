package builder

import (
	"go.uber.org/zap"
	"service-worker-sqs-postgres/core/domain"
	"service-worker-sqs-postgres/dataproviders/processor"
	"service-worker-sqs-postgres/dataproviders/repository/events"
)

// NewProcessor define all usecases to be instantiated Processor associated with the consumer.
func NewProcessor(logger *zap.SugaredLogger, source domain.Source, er *events.EventsRepository) (*processor.Processor, error) {
	return processor.New(logger, source, er)
}
