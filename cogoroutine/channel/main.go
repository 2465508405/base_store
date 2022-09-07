/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-06 15:10:43
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-06 16:40:31
 * @FilePath: /allfunc/cogoroutine/channel/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE

 */
package main

import "fmt"

func main() {

	ch4 := make(chan int)

	go Rece1(ch4)
	go Rece2(ch4)
	func() {
		for i := 1; i < 100; i++ {
			ch4 <- i
		}
		close(ch4)
	}()

	//接收方法1
	// for v := range ch4 { // 通道关闭后会退出for range循环

	// 	fmt.Println(v)
	// }

	//接收方法2
	// for {
	// 	if v, ok := <-ch4; ok {
	// 		fmt.Println(v)
	// 	} else {
	// 		break
	// 	}
	// }

}

func Rece1(ch chan int) {
	for v := range ch { // 通道关闭后会退出for range循环

		fmt.Println(v)
	}
}

func Rece2(ch chan int) {
	for v := range ch { // 通道关闭后会退出for range循环

		fmt.Println(v)
	}
}
