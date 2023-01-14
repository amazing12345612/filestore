package main

import (
	"log"

	micro "github.com/asim/go-micro/v3"

	"filestore/service/account/handler"
	proto "filestore/service/account/proto"
	"github.com/asim/go-micro/plugins/registry/consul/v3"
	"github.com/asim/go-micro/v3/registry"
)

const (
	ServerName = "go.micro.api.user"
	ConsulAddr = "127.0.0.1:8500"
)

func main() {
	// 创建一个service
	reg := consul.NewRegistry(registry.Addrs(ConsulAddr))
	service := micro.NewService(
		micro.Name(ServerName),
		micro.Address("0.0.0.0:8001"),
		// micro.RegisterTTL(time.Second*10),
		// micro.RegisterInterval(time.Second*5),
		micro.Registry(reg),
	)
	service.Init()

	proto.RegisterUserServiceHandler(service.Server(), new(handler.User))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
