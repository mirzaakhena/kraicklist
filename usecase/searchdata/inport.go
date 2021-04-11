package searchdata

import (
	"challenge.haraj.com.sa/kraicklist/domain/entity"
	"context"
)

// Inport of SearchData
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase SearchData
type InportRequest struct {
	Keyword string
}

// InportResponse is response payload after running the usecase SearchData
type InportResponse struct {
	Result           []*entity.RawData
	ContainedKeyword []string
}
