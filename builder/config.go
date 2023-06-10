package builder

import env "service-template-golang/utils"

type Configuration struct {
	Port                 int
	ApplicationID        string
	LogLevel             string
	Region               string
	AccessKey            string
	SecretKey            string
	SQSUrl               string
	SQSMaxMessages       int
	SQSVisibilityTimeout int
}

func LoadConfig() (*Configuration, error) {
	applicationID, err := env.GetString("APPLICATION_ID")
	if err != nil {
		return nil, err
	}

	port, err := env.GetInt("SERVER_PORT")
	if err != nil {
		return nil, err
	}

	loglevel, err := env.GetString("LOG_LEVEL")
	if err != nil {
		return nil, err
	}

	access, err := env.GetString("AWS_ACCESS_KEY")
	if err != nil {
		return nil, err
	}

	secret, err := env.GetString("AWS_SECRET_KEY")
	if err != nil {
		return nil, err
	}

	region, err := env.GetString("AWS_REGION")
	if err != nil {
		return nil, err
	}

	sqsUrl, err := env.GetString("AWS_SQS_URL")
	if err != nil {
		return nil, err
	}

	sqsMaxMessages, err := env.GetInt("AWS_SQS_MAX_MESSAGES")
	if err != nil {
		return nil, err
	}

	sqsVisibilityTimeout, err := env.GetInt("AWS_SQS_VISIBILITY_TIMEOUT")
	if err != nil {
		return nil, err
	}

	return &Configuration{
		Port:                 port,
		ApplicationID:        applicationID,
		LogLevel:             loglevel,
		AccessKey:            access,
		SecretKey:            secret,
		Region:               region,
		SQSUrl:               sqsUrl,
		SQSMaxMessages:       sqsMaxMessages,
		SQSVisibilityTimeout: sqsVisibilityTimeout,
	}, nil
}
