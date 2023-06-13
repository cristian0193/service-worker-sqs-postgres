package builder

import (
	"go.uber.org/zap"
	"service-template-golang/database"
	"service-template-golang/domain"
	"service-template-golang/processor"
)

// NewProcessor define all services to be instantiated Processor associated with the consumer.
func NewProcessor(logger *zap.SugaredLogger, source domain.Source, db *database.ClientDB) (*processor.Processor, error) {
	return processor.New(logger, source, db)
}
