/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-06-11 10:59:48
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-06-11 14:49:55
 * @FilePath: /allfunc/conn_pools/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"context"
	"fmt"

	"project/allfunc/conn_pools/grpcpool"

	"google.golang.org/grpc"
)

func main() {
	p, err := grpcpool.New(func() (*grpc.ClientConn, error) {
		return grpc.Dial("127.0.0.1", grpc.WithInsecure())
	}, 1, 3, 0)
	if err != nil {
		fmt.Printf("The pool returned an error: %s\n", err.Error())
	}

	// Get a client
	client, err := p.Get(context.Background())
	fmt.Println("777")
	if err != nil {
		fmt.Printf("Get returned an error: %s\n", err.Error())
	}
	client.Close()
	if client == nil {
		fmt.Println("client was nil")
	} else {
		fmt.Printf("%+v\n", client)
	}
	client1, err := p.Get(context.Background())
	if err != nil {
		fmt.Printf("Get returned an error: %s\n", err.Error())
	}
	if client1 == nil {
		fmt.Println("client was nil")
	} else {
		fmt.Printf("%+v\n", client1)
	}
	client1.Close()
	client2, err := p.Get(context.Background())
	if err != nil {
		fmt.Printf("Get returned an error: %s\n", err.Error())
	}
	if client2 == nil {
		fmt.Println("client was nil")
	} else {
		fmt.Printf("%+v\n", client2)
	}

	client2.Close()
	client3, err := p.Get(context.Background())
	if err != nil {
		fmt.Printf("Get returned an error: %s\n", err.Error())
	}
	if client3 == nil {
		fmt.Println("client was nil")
	} else {
		fmt.Printf("%+v\n", client3)
	}
	client3.Close()
	client4, err := p.Get(context.Background())
	if err != nil {
		fmt.Printf("Get returned an error: %s\n", err.Error())
	}
	if client4 == nil {
		fmt.Println("client was nil")
	} else {
		fmt.Printf("%+v\n", client4)
	}

	client4.Close()
}
