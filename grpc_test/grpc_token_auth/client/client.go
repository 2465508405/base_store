/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-02 17:14:44
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-07-03 17:10:51
 * @FilePath: /allfunc/grpc_test/client/client.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"context"
	"fmt"
	"project/allfunc/grpc_test/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type Credentials struct {
}

func (per Credentials) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "1001",
		"appkey": "hobby",
	}, nil
}

// RequireTransportSecurity indicates whether the credentials requires
// transport security.
func (per Credentials) RequireTransportSecurity() bool {
	return false
}
func main() {

	// inter := func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// 	start := time.Now()
	// 	md := metadata.New(map[string]string{
	// 		"appid":  "10001",
	// 		"appkey": "hobby",
	// 	})
	// 	ctx = metadata.NewOutgoingContext(context.Background(), md)
	// 	err := invoker(ctx, method, req, reply, cc, opts...)
	// 	fmt.Printf("耗时:%s", time.Since(start))
	// 	fmt.Println("\n")
	// 	return err
	// }
	// var cred = PerRPCCredentials{}
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithInsecure())
	opts = append(opts, grpc.WithPerRPCCredentials(Credentials{}))
	// opts = append(opts, grpc.WithUnaryInterceptor(inter))

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
