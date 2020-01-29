package main

import (
	"fmt"

	"github.com/kytra-app/portfolio-allocation-srv/handler"
	proto "github.com/kytra-app/portfolio-allocation-srv/proto"
	"github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/kubernetes"
)

func main() {
	// Create The Service
	service := micro.NewService(
		micro.Name("kytra-srv-v1-portfolio-allocation"),
		micro.Version("latest"),
	)
	service.Init()

	// Register to Service Discovery
	hander := handler.New(service.Client())
	proto.RegisterPortfolioAllocationHandler(service.Server(), hander)

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}