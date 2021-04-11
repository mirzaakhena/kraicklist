package loaddata

import (
	"challenge.haraj.com.sa/kraicklist/domain/repository"
	"challenge.haraj.com.sa/kraicklist/domain/service"
)

// Outport of LoadData
type Outport interface {
	repository.TransactionRepo
	service.ReadJSONDataService
	repository.SaveDataRepo
}
