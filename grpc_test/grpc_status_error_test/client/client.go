/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-02 17:14:44
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-07-03 20:04:44
 * @FilePath: /allfunc/grpc_test/client/client.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"context"
	"fmt"
	"project/allfunc/grpc_test/grpc_status_error_test/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	// md := metadata.Pairs("timestamp", time.Now().Format(timestamp))
	//超时机制
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)
	_, err = c.SayHello(ctx, &proto.HelloRequest{Name: "chenshimei"})
	if err != nil {

		st, ok := status.FromError(err)
		if !ok {
			panic(err)
		}
		code := st.Code()
		msg := st.Message()
		fmt.Println(msg)
		fmt.Println(code)

	}

	// fmt.Printf("%s", r.Message)
}
