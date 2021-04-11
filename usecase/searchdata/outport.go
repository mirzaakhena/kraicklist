package searchdata

import "challenge.haraj.com.sa/kraicklist/domain/repository"

// Outport of SearchData
type Outport interface {
	repository.FindDataByKeywordRepo
}
