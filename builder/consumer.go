package builder

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"go.uber.org/zap"
	"service-template-golang/clients/services"
	"service-template-golang/consumer"
	"service-template-golang/domain"
	env "service-template-golang/utils"
)

func NewSQSConsumer(logger *zap.SugaredLogger) (domain.Source, error) {
	region, err := env.GetString("AWS_REGION")
	if err != nil {
		return nil, err
	}

	dqlqURL, err := env.GetString("DELTA_SQS_URL")
	if err != nil {
		return nil, err
	}

	logger.Infof("dead letter queue SQS url [%s] Region [%s]", dqlqURL, region)

	sqsMaxMessages, err := env.GetInt("DELTA_SQS_MAX_MESSAGES")
	if err != nil {
		return nil, err
	}

	sqsVisibilityTimeout, err := env.GetInt("DELTA_SQS_VISIBILITY_TIMEOUT")
	if err != nil {
		return nil, err
	}

	sqsSessionConfig := &aws.Config{Region: aws.String(region)}
	sqsSessionConfig.WithCredentials(credentials.NewStaticCredentials("id", "temp", "temp"))
	sqsSessionConfig.WithEndpoint("http://localhost:4576")
	sqsSession := session.Must(session.NewSession(sqsSessionConfig))

	sqs, err := services.NewSQSClient(sqsSession, dqlqURL, sqsMaxMessages, sqsVisibilityTimeout)
	if err != nil {
		return nil, fmt.Errorf("sqsclient.NewWithSession: %w", err)
	}

	source, err := consumer.New(sqs, logger, 5)
	if err != nil {
		return nil, fmt.Errorf("sqssource.New: %w", err)
	}

	return source, nil
}
