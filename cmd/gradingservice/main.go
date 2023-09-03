package main

import (
	"context"
	"fmt"

	stdlog "log"

	"github.com/xunfei/microservice/grades"
	"github.com/xunfei/microservice/log"
	"github.com/xunfei/microservice/registry"
	"github.com/xunfei/microservice/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)
	r := registry.Registration{
		ServiceName:       registry.GradingService,
		ServiceURL:        serviceAddress,
		RequiredSerivices: []registry.ServiceName{registry.LogService},
		ServiceUpdateURL:  serviceAddress + "/serivices",
	}
	ctx, err := service.Start(context.Background(),
		r,
		host,
		port,
		grades.RegisterHandlers,
	)
	if err != nil {
		stdlog.Fatal(err)
	}
	if logProvider,err:= registry.GetProvider(registry.LogService);err!=nil{
		fmt.Printf("Logging service found at: %s\n",logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}
	<-ctx.Done()
	fmt.Println("Shutting down grading service")
}
