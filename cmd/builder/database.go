package builder

import (
	"service-template-golang/database"
)

func NewDB(config *Configuration) (*database.ClientDB, error) {
	db := database.NewDBClient(config.DBHost, config.DBUsername, config.DBPassword, config.DBName, config.DBPort)
	err := db.Open()
	return db, err
}
