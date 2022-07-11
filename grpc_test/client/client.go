/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-02 17:14:44
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-07-02 17:18:17
 * @FilePath: /allfunc/grpc_test/client/client.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"context"
	"fmt"
	"project/allfunc/grpc_test/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)
	r, err := c.SayHello(context.Background(), &proto.HelloRequest{Name: "chenshimei"})
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", r.Message)
}
