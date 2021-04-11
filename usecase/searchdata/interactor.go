package searchdata

import (
	"challenge.haraj.com.sa/kraicklist/application/apperror"
	"context"
)

//go:generate mockery --dir port/ --name SearchDataOutport -output mocks/

type searchDataInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase SearchData
func NewUsecase(outputPort Outport) Inport {
	return &searchDataInteractor{
		outport: outputPort,
	}
}

// Execute the usecase SearchData
func (r *searchDataInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	rawDataObjs, err := r.outport.FindDataByKeyword(ctx, req.Keyword)
	if err != nil {
		return nil, err
	}
	if rawDataObjs == nil {
		return nil, apperror.ObjectNotFound
	}

	res.ContainedKeyword = rawDataObjs.ContainedKeyword
	res.Result = rawDataObjs.Result

	return res, nil
}
