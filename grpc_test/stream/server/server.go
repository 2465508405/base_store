/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-02 18:32:44
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-07-02 20:32:22
 * @FilePath: /allfunc/grpc_test/stream/server/server.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"net"
	"project/allfunc/grpc_test/stream/proto"
	"sync"
	"time"

	"google.golang.org/grpc"
)

const PORT = "50051"

type Server struct {
	proto.UnimplementedGreeterServer
}

func (s *Server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {

	i := 0
	for {
		i++

		if i < 10 {
			err := res.Send(&proto.StreamResData{
				Data: fmt.Sprintf("current:%v", time.Now().Unix()),
			})
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(time.Second)
		} else {
			break
		}

	}
	return nil
}

func (s *Server) PutStream(cliStr proto.Greeter_PutStreamServer) error {
	for {
		req, err := cliStr.Recv()
		if err != nil {
			panic(err)
		}

		fmt.Println(req.Data)
	}

	return nil
}

func (s *Server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			req, err := allStr.Recv()
			if err != nil {
				break
			}
			fmt.Println(req.Data)
		}

	}()
	go func() {
		defer wg.Done()
		for {
			err := allStr.Send(&proto.StreamResData{Data: "hello client"})
			if err != nil {
				break
			}
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()

	return nil
}

func main() {
	g := grpc.NewServer()
	proto.RegisterGreeterServer(g, &Server{})
	lis, err := net.Listen("tcp", "127.0.0.1:"+PORT)
	if err != nil {
		panic(err)
	}
	err = g.Serve(lis)
	if err != nil {
		panic("start err:" + err.Error())
	}

}
