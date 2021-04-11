package userapi

import (
	"net/http"

	"challenge.haraj.com.sa/kraicklist/infrastructure/log"
	"challenge.haraj.com.sa/kraicklist/infrastructure/util"
	"challenge.haraj.com.sa/kraicklist/usecase/searchdata"
	"github.com/gin-gonic/gin"
)

// searchDataHandler ...
func (r *Controller) searchDataHandler(inputPort searchdata.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		ctx := log.ContextWithLogGroupID(c.Request.Context())

		var req searchdata.InportRequest
		req.Keyword = c.DefaultQuery("q", "")

		if req.Keyword == "" {
			c.JSON(http.StatusOK, NewSuccessResponse("not found"))
			return
		}

		log.InfoRequest(ctx, util.MustJSON(req))

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			log.ErrorResponse(ctx, err)
			c.JSON(http.StatusBadRequest, NewErrorResponse(err))
			return
		}

		log.InfoResponse(ctx, util.MustJSON("Result Found"))
		c.JSON(http.StatusOK, NewSuccessResponse(res))

	}
}
