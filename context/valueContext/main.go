/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-28 14:31:31
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-28 14:34:43
 * @FilePath: /allfunc/context/valueContext/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"context"
	"fmt"
	"time"
)

func HandelRequest(ctx context.Context) {
	go ExecuteMysql(ctx)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running, parameter: ", ctx.Value("parameter"))
			time.Sleep(2 * time.Second)
		}
	}
}

func ExecuteMysql(ctx context.Context) {
	fmt.Println("mysql:paramter=>", ctx.Value("parameter"))
	for {
		select {
		case <-ctx.Done():
			fmt.Println("HandelRequest Done.")
			return
		default:
			fmt.Println("HandelRequest running, parameter: ", ctx.Value("parameter"))
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	ctx := context.WithValue(context.Background(), "parameter", "1")
	go HandelRequest(ctx)

	time.Sleep(10 * time.Second)
}
