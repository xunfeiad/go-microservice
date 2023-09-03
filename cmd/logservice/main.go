package main

import (
	"context"
	"fmt"

	"github.com/xunfei/microservice/log"
	"github.com/xunfei/microservice/registry"

	stdlog "log"

	"github.com/xunfei/microservice/service"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)
	reg := registry.Registration{
		ServiceName:       registry.LogService,
		ServiceURL:        serviceAddress,
		RequiredSerivices: make([]registry.ServiceName, 0),
		ServiceUpdateURL:  serviceAddress + "/services",
	}
	ctx, err := service.Start(
		context.Background(),
		reg,
		host,
		port,
		log.RegisterHandlers,
	)
	if err != nil {
		stdlog.Fatalln(err)
	}
	<-ctx.Done()

	fmt.Println("Shutting down log service.")
}
