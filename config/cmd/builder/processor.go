package builder

import (
	"go.uber.org/zap"
	"service-worker-sqs-postgres/core/domain"
	"service-worker-sqs-postgres/dataproviders/database"
	"service-worker-sqs-postgres/dataproviders/processor"
)

// NewProcessor define all usecases to be instantiated Processor associated with the consumer.
func NewProcessor(logger *zap.SugaredLogger, source domain.Source, db *database.ClientDB) (*processor.Processor, error) {
	return processor.New(logger, source, db)
}
