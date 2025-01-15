package dbservice

import (
	"github.com/romangorisek/gpttui/logger"
)

type SqliteService struct {
	ConnectionString string
}

func NewSqliteService(connectionString string) *SqliteService {
	return &SqliteService{ConnectionString: connectionString}
}

func (db SqliteService) FetchData() (string, error) {
	logger.Log.Info("test")
	data := ""
	return data, nil
}
