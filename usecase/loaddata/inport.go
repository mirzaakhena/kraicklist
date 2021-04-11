package loaddata

import (
	"context"
)

// Inport of LoadData
type Inport interface {
	Execute(ctx context.Context, req InportRequest) (*InportResponse, error)
}

// InportRequest is request payload to run the usecase LoadData
type InportRequest struct {
	Filename string
}

// InportResponse is response payload after running the usecase LoadData
type InportResponse struct {
}
