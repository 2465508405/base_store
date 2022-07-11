/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-02 16:59:11
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-07-03 20:05:06
 * @FilePath: /allfunc/grpc_test/server/server.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"context"
	"net"
	"project/allfunc/grpc_test/grpc_status_error_test/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	time.Sleep(time.Second * 6)
	err := status.Error(codes.NotFound, "invalid username ")
	return &proto.HelloReply{Message: "hello :" + req.Name}, err
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	err = g.Serve(lis)
	if err != nil {
		panic("start fail" + err.Error())
	}
}
