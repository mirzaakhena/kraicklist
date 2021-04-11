package userapi

import (
	"challenge.haraj.com.sa/kraicklist/infrastructure/log"
	"challenge.haraj.com.sa/kraicklist/infrastructure/util"
	"challenge.haraj.com.sa/kraicklist/usecase/loaddata"
	"context"
)

// loadDataHandler ...
func (r *Controller) loadDataHandler(inputPort loaddata.Inport) {

	ctx := log.ContextWithLogGroupID(context.Background())

	var req loaddata.InportRequest
	req.Filename = "data.gz"

	log.InfoRequest(ctx, util.MustJSON(req))

	res, err := inputPort.Execute(ctx, req)
	if err != nil {
		log.ErrorResponse(ctx, err)
		return
	}

	log.InfoResponse(ctx, util.MustJSON(res))

}
