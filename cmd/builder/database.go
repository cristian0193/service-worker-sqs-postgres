package builder

import (
	"service-worker-sqs-postgres/database"
)

// NewDB defines all configurations to instantiate a database client.
func NewDB(config *Configuration) (*database.ClientDB, error) {
	db := database.NewDBClient(config.DBHost, config.DBUsername, config.DBPassword, config.DBName, config.DBPort)
	err := db.Open()

	return db, err
}
