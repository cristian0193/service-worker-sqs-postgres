package builder

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"go.uber.org/zap"
	"service-template-golang/clients/awssqs"
	"service-template-golang/consumer"
	"service-template-golang/database"
	"service-template-golang/domain"
)

// NewSQS define all services to instantiate SQS.
func NewSQS(logger *zap.SugaredLogger, config *Configuration, sessionaws *session.Session, db *database.ClientDB) (domain.Source, error) {
	sqs, err := awssqs.NewSQSClient(sessionaws, config.SQSUrl, config.SQSMaxMessages, config.SQSVisibilityTimeout)
	if err != nil {
		return nil, fmt.Errorf("error awssqs.NewSQSClient: %w", err)
	}

	source, err := consumer.New(sqs, logger, config.SQSMaxMessages, db)
	if err != nil {
		return nil, fmt.Errorf("error consumer.New: %w", err)
	}

	return source, nil
}
