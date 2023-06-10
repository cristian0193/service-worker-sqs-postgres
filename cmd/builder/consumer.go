package builder

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"go.uber.org/zap"
	"service-template-golang/clients/awssqs"
	"service-template-golang/consumer"
	"service-template-golang/domain"
)

func NewSQSConsumer(logger *zap.SugaredLogger, config *Configuration) (domain.Source, error) {

	sqsSessionConfig := &aws.Config{
		Region:      aws.String(config.Region),
		Endpoint:    aws.String(config.SQSUrl),
		MaxRetries:  aws.Int(3),
		Credentials: credentials.NewStaticCredentials(config.AccessKey, config.SecretKey, ""),
	}

	sqsSession := session.Must(session.NewSession(sqsSessionConfig))

	sqs, err := awssqs.NewSQSClient(sqsSession, config.SQSUrl, config.SQSMaxMessages, config.SQSVisibilityTimeout)
	if err != nil {
		return nil, fmt.Errorf("error awssqs.NewSQSClient: %w", err)
	}

	source, err := consumer.New(sqs, logger, config.SQSMaxMessages)
	if err != nil {
		return nil, fmt.Errorf("error consumer.New: %w", err)
	}

	return source, nil
}
