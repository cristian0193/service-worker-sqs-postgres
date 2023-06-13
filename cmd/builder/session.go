package builder

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

// NewSession define all configuration to instantiate a session aws.
func NewSession(config *Configuration) (*session.Session, error) {
	sqsSessionConfig := &aws.Config{
		Region:      aws.String(config.Region),
		Endpoint:    aws.String(config.SQSUrl),
		MaxRetries:  aws.Int(3),
		Credentials: credentials.NewStaticCredentials(config.AccessKey, config.SecretKey, ""),
	}
	sess, err := session.NewSession(sqsSessionConfig)
	if err != nil {
		return nil, err
	}

	return session.Must(sess, err), nil
}
