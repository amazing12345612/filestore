package main

import (
	dbProxy "filestore/service/dbproxy/proto"
	dbRpc "filestore/service/dbproxy/rpc"
	"log"

	"github.com/asim/go-micro/plugins/registry/consul/v3"
	micro "github.com/asim/go-micro/v3"
	"github.com/asim/go-micro/v3/registry"
)

var (
	ConsulAddr = "127.0.0.1:8500"
	ServerName = "go.micro.service.dbproxy"
)

func startRpcService() {
	reg := consul.NewRegistry(registry.Addrs(ConsulAddr))
	service := micro.NewService(
		micro.Name(ServerName),
		micro.Address("0.0.0.0:8001"),
		// micro.RegisterTTL(time.Second*10),
		// micro.RegisterInterval(time.Second*5),
		micro.Registry(reg),
	)
	service.Init()
	dbProxy.RegisterDBProxyServiceHandler(service.Server(), new(dbRpc.DBProxy))
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}

func main() {
	startRpcService()
	// res, err := mapper.FuncCall("/user/UserExist", []interface{}{"haha"}...)
	// log.Printf("error: %+v\n", err)
	// log.Printf("result: %+v\n", res[0].Interface())

	// res, err = mapper.FuncCall("/user/UserExist", []interface{}{"admin"}...)
	// log.Printf("error: %+v\n", err)
	// log.Printf("result: %+v\n", res[0].Interface())
}
