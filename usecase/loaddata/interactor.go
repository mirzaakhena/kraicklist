package loaddata

import (
	"challenge.haraj.com.sa/kraicklist/domain/entity"
	"challenge.haraj.com.sa/kraicklist/domain/service"
	"context"
)

//go:generate mockery --dir port/ --name LoadDataOutport -output mocks/

type loadDataInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase LoadData
func NewUsecase(outputPort Outport) Inport {
	return &loadDataInteractor{
		outport: outputPort,
	}
}

// Execute the usecase LoadData
func (r *loadDataInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	// use transaction here
	err := service.ExecuteTransaction(ctx, r.outport, func(dbCtx context.Context) error {

		// read data from file
		err := r.outport.ReadJSONData(ctx, service.ReadJSONDataServiceRequest{
			Filename: req.Filename,
			ReadDataPerline: func(data entity.RawData) error {
				err := r.outport.SaveData(ctx, &data)
				if err != nil {
					return err
				}
				return nil
			},
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return res, nil
}
