package main

import (
	"context"
	pb "filestore/service/user/proto" // 引入proto包
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:12344"
)

func main() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}
	defer conn.Close()

	// 初始化客户端
	c := pb.NewUserServiceClient(conn)

	// 调用方法
	req := &pb.ReqSignup{Username: "admin111", Password: "admin111"}
	res, err := c.Signup(context.Background(), req)

	if err != nil {
		grpclog.Fatalln(err)
	}

	fmt.Println(res.GetMessage())
}
