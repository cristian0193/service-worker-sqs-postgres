package builder

import (
	"go.uber.org/zap"
	"service-template-golang/domain"
	"service-template-golang/processor"
)

func NewProcessor(logger *zap.SugaredLogger, source domain.Source) (*processor.Processor, error) {
	return processor.New(logger, source)
}
