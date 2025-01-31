package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// GinHTTPHandler will define basic HTTP configuration with gracefully shutdown
type GinHTTPHandler struct {
	GracefullyShutdown
	Router *gin.Engine
}

func NewGinHTTPHandler(address string) GinHTTPHandler {

	router := gin.Default()

	router.Static("./static", "/static")
	router.LoadHTMLGlob("static/index.html")

	// PING API
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Ready")
	})

	// CORS
	router.Use(cors.New(cors.Config{
		ExposeHeaders:   []string{"Data-Length"},
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "PATCH", "HEAD", "OPTIONS"},
		AllowAllOrigins: true,
		AllowHeaders:    []string{"Content-Type", "Authorization"},
		MaxAge:          12 * time.Hour,
	}))

	return GinHTTPHandler{
		GracefullyShutdown: NewGracefullyShutdown(router, address),
		Router:             router,
	}

}

// RunApplication is implementation of RegistryContract.RunApplication()
func (r *GinHTTPHandler) RunApplication() {
	r.RunWithGracefullyShutdown()
}
