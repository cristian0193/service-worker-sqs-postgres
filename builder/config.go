package builder

import env "service-template-golang/utils"

type Configuration struct {
	Port          int
	ApplicationID string
	LogLevel      string
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

	return &Configuration{
		Port:          port,
		ApplicationID: applicationID,
		LogLevel:      loglevel,
	}, nil
}
