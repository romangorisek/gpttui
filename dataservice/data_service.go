package dataservice

import (
	"github.com/romangorisek/gpttui/dbservice"
	"github.com/romangorisek/gpttui/networkservice"
)

type DataService struct {
	dbService      dbservice.DbService
	networkService networkservice.NetworkService
}

func NewDataService(dbService dbservice.DbService, netnetworkService networkservice.NetworkService) *DataService {
	return &DataService{
		dbService:      dbService,
		networkService: netnetworkService,
	}
}
