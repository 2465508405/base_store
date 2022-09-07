/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-06 14:03:14
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-06 14:42:15
 * @FilePath: /allfunc/cogoroutine/sync/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Hello(i int) {
	defer wg.Done()
	fmt.Println("print:", i)
}

func main() {

	// 合起来写
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}
