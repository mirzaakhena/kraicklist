package main

import (
	"flag"

	"challenge.haraj.com.sa/kraicklist/application"
	"challenge.haraj.com.sa/kraicklist/application/registry"
)

func main() {
	flag.Parse()
	switch flag.Arg(0) {
	default:
		application.Run(registry.NewApp1())
	}
}
