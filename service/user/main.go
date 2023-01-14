package main

import (
	"context"
	"filestore/common"
	cfg "filestore/config"
	mydb "filestore/db"
	client "filestore/service/dbproxy/client"
	pb "filestore/service/user/proto"
	"filestore/util"
	"fmt"
	"net"
	"time"

	"google.golang.org/grpc"
)

type User struct {
	pb.UnimplementedUserServiceServer
}

// GenToken : 生成token
func GenToken(username string) string {
	// 40位字符:md5(username+timestamp+token_salt)+timestamp[:8]
	ts := fmt.Sprintf("%x", time.Now().Unix())
	tokenPrefix := util.MD5([]byte(username + ts + "_tokensalt"))
	return tokenPrefix + ts[:8]
}

// 注册接口
func (u *User) Signup(ctx context.Context, req *pb.ReqSignup) (*pb.RespSignup, error) {
	res := new(pb.RespSignup)
	username := req.Username
	passwd := req.Password

	// 参数简单校验
	if len(username) < 3 || len(passwd) < 5 {
		res.Code = common.StatusParamInvalid
		res.Message = "注册参数无效"
		return res, nil
	}

	// 对密码进行加盐及取Sha1值加密
	encPasswd := util.Sha1([]byte(passwd + cfg.PwdSalt))
	// 将用户信息注册到用户表中

	//TODO修改成微服务的方式
	res, err := client.UserSignup(username, encPasswd)
	if err != nil {
		res.Code = common.StatusRegisterFailed
		res.Message = "注册失败"
		return res, err
	}
	res.Code = common.StatusOK
	res.Message = "注册成功"
	return res, nil
}

// Signin : 处理用户登录请求
func (u *User) Signin(ctx context.Context, req *pb.ReqSignin) (*pb.RespSignin, error) {
	res := new(pb.RespSignin)
	username := req.Username
	password := req.Password

	encPasswd := util.Sha1([]byte(password + cfg.PwdSalt))

	// 1. 校验用户名及密码
	dbResp := mydb.UserSignin(username, encPasswd)
	if dbResp != true {
		res.Code = common.StatusLoginFailed
		return res, nil
	}
	fmt.Println("密码正确")

	// 2. 生成访问凭证(token)
	token := GenToken(username)
	upRes := mydb.UpdateToken(username, token)
	if upRes != true {
		res.Code = common.StatusServerError
		return res, nil
	}

	// 3. 登录成功, 返回token
	res.Code = common.StatusOK
	res.Token = token
	return res, nil
}

func main() {
	// 创建一个service
	listen, err := net.Listen("tcp", "127.0.0.1:12344")
	if err != nil {
		fmt.Println("12121231")
	}
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, new(User))
	server.Serve(listen)

}
