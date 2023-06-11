package builder

import (
	"go.uber.org/zap"
	"service-template-golang/database"
	"service-template-golang/domain"
	"service-template-golang/processor"
)

func NewProcessor(logger *zap.SugaredLogger, source domain.Source, rds *database.ClientDB) (*processor.Processor, error) {
	return processor.New(logger, source, rds)
}
