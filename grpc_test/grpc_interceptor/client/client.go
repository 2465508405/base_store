/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-02 17:14:44
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-07-03 16:19:32
 * @FilePath: /allfunc/grpc_test/client/client.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"context"
	"fmt"
	"project/allfunc/grpc_test/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {

	inter := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		start := time.Now()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fmt.Printf("耗时:%s", time.Since(start))

		return err
	}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithUnaryInterceptor(inter))

	// conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure(), opt)
	conn, err := grpc.Dial("127.0.0.1:50051", opts...)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	// md := metadata.Pairs("timestamp", time.Now().Format(timestamp))
	md := metadata.New(map[string]string{
		"name": "bobby",
		"url":  "imooc.com",
	})
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: "chenshimei"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", r.Message)
}
