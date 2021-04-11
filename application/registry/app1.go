package registry

import (
	"challenge.haraj.com.sa/kraicklist/application"
	"challenge.haraj.com.sa/kraicklist/controller/userapi"
	"challenge.haraj.com.sa/kraicklist/gateway"
	"challenge.haraj.com.sa/kraicklist/infrastructure/server"
	"challenge.haraj.com.sa/kraicklist/usecase/loaddata"
	"challenge.haraj.com.sa/kraicklist/usecase/searchdata"
	"fmt"
	"os"
)

type app1 struct {
	server.GinHTTPHandler
	userapiController userapi.Controller
	// TODO Another controller will added here ... <<<<<<
}

func NewApp1() application.RegistryContract {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3001"
	}

	address := fmt.Sprintf(":%s", port)
	httpHandler := server.NewGinHTTPHandler(address)
	datasource := gateway.NewInmemoryGateway()

	return &app1{
		GinHTTPHandler: httpHandler,
		userapiController: userapi.Controller{
			Router:         httpHandler.Router,
			LoadDataInport: loaddata.NewUsecase(datasource),
			SearchDataInport: searchdata.NewUsecase(datasource),
			// TODO another Inport will added here ... <<<<<<
		},
		// TODO another controller will added here ... <<<<<<
	}
}

func (r *app1) SetupController() {
	r.userapiController.RegisterRouter()
	// TODO another router call will added here ... <<<<<<
}
