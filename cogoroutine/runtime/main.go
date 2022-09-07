package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

func main() {

	runtime.GOMAXPROCS(2)
	go a()
	go b()

	time.Sleep(time.Second * 3)
	// go func(s string) {
	// 	for i := 0; i < 2; i++ {
	// 		fmt.Println(s)
	// 	}
	// }("world")
	// // 主协程
	// for i := 0; i < 2; i++ {
	// 	// 切一下，再次分配任务
	// 	runtime.Gosched() //让出CPU时间片，重新等待安排任务
	// 	fmt.Println("hello")
	// }
}
