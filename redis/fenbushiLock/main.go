/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-05-28 14:36:34
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-05-28 15:09:36
 * @FilePath: /allfunc/redis/fenbushiLock/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"github.com/stvp/tempredis"
	// "github.com/stvp/tempredis"
)

func main() {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	fmt.Println(base64.StdEncoding.EncodeToString(b))

	ctx := context.Background()

	fmt.Println(ctx)
	server, err := tempredis.Start(tempredis.Config{})
	if err != nil {
		panic(err)
	}
	defer server.Term()

	client83 := goredislib.NewClient(&goredislib.Options{
		Addr: "localhost:6383",
	})
	client84 := goredislib.NewClient(&goredislib.Options{
		Addr: "localhost:6384",
	})
	client85 := goredislib.NewClient(&goredislib.Options{
		Addr: "localhost:6385",
	})
	client86 := goredislib.NewClient(&goredislib.Options{
		Addr: "localhost:6386",
	})

	pool80 := goredis.NewPool(client83)
	pool81 := goredis.NewPool(client84)
	pool82 := goredis.NewPool(client85)
	pool83 := goredis.NewPool(client86)

	rs := redsync.New(pool80, pool81, pool82, pool83)

	mutex := rs.NewMutex("test-redsync")
	// ctx := context.Background()

	if err := mutex.Lock(); err != nil {
		panic(err)
	} else {

		fmt.Println("locked ")
	}
	fn := make(chan int)
	// c := make(chan int)
	go func(mutex *redsync.Mutex, fn chan int) {

		for {
			time.Sleep(time.Second * 3)
			select {
			case <-fn:
				return
			default:
				mutex.Extend()
			}
		}

	}(mutex, fn)
	time.Sleep(time.Second * 20)
	fn <- 1
	if _, err := mutex.Unlock(); err != nil {
		panic(err)
	}
	fmt.Println("unlock success")
}
