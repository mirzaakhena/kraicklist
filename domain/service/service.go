package service

import (
	"challenge.haraj.com.sa/kraicklist/domain/entity"
	"context"
)

type ReadJSONDataService interface {
	ReadJSONData(ctx context.Context, req ReadJSONDataServiceRequest) error
}

type ReadJSONDataServiceRequest struct {
	Filename        string
	ReadDataPerline func(data entity.RawData) error
}
