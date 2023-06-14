package builder

import (
	"go.uber.org/zap"
	"service-worker-sqs-postgres/database"
	"service-worker-sqs-postgres/domain"
	"service-worker-sqs-postgres/processor"
)

// NewProcessor define all services to be instantiated Processor associated with the consumer.
func NewProcessor(logger *zap.SugaredLogger, source domain.Source, db *database.ClientDB) (*processor.Processor, error) {
	return processor.New(logger, source, db)
}
