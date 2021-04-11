package userapi

import (
	"challenge.haraj.com.sa/kraicklist/usecase/loaddata"
	"challenge.haraj.com.sa/kraicklist/usecase/searchdata"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Controller struct {
	Router           gin.IRouter
	LoadDataInport   loaddata.Inport
	SearchDataInport searchdata.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {

	r.Router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "MainMenu"})
	})

	r.loadDataHandler(r.LoadDataInport)

	r.Router.GET("/search", r.authorized(), r.searchDataHandler(r.SearchDataInport))
}
