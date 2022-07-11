/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-02 19:10:45
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-07-02 20:37:08
 * @FilePath: /allfunc/grpc_test/stream/client/client.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"context"
	"fmt"
	"project/allfunc/grpc_test/stream/proto"
	"sync"
	"time"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	cli := proto.NewGreeterClient(conn)

	// res, err := cli.GetStream(context.Background(), &proto.StreamReqData{Data: "张三"})
	// if err != nil {
	// 	panic(err)
	// }
	// for {
	// 	recv, err := res.Recv()
	// 	if err != nil {
	// 		break
	// 	}
	// 	fmt.Println(recv.Data)
	// }
	// put, err := cli.PutStream(context.Background())
	// if err != nil {
	// 	panic(err)
	// }
	// for {
	// 	err := put.Send(&proto.StreamReqData{
	// 		Data: "hello server",
	// 	})
	// 	if err != nil {
	// 		break
	// 	}
	// 	time.Sleep(time.Second)
	// }

	all, err := cli.AllStream(context.Background())
	if err != nil {
		panic(err)
	}
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			err = all.Send(&proto.StreamReqData{Data: "hello server"})
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Second)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			rec, err := all.Recv()
			if err != nil {
				panic(err)
			}
			fmt.Println(rec.Data)
		}
	}()
	wg.Wait()
	time.Sleep(time.Second * 100)

}
