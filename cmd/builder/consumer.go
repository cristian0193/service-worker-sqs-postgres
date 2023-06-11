package builder

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"go.uber.org/zap"
	"service-template-golang/clients/awssqs"
	"service-template-golang/consumer"
	"service-template-golang/domain"
)

func NewSQS(logger *zap.SugaredLogger, config *Configuration, sessionaws *session.Session) (domain.Source, error) {
	sqs, err := awssqs.NewSQSClient(sessionaws, config.SQSUrl, config.SQSMaxMessages, config.SQSVisibilityTimeout)
	if err != nil {
		return nil, fmt.Errorf("error awssqs.NewSQSClient: %w", err)
	}

	source, err := consumer.New(sqs, logger, config.SQSMaxMessages)
	if err != nil {
		return nil, fmt.Errorf("error consumer.New: %w", err)
	}

	return source, nil
}
