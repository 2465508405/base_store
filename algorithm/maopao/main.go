/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-10 16:49:47
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-07-11 16:04:42
 * @FilePath: /allfunc/algorithm/maopao/main.go
 */
package main

import (
	"fmt"
)

func main() {

	// var arr []int = make([]int, 3)
	arr := []int{1, 232, 23, 2223, 2, 8, 10}

	for end := len(arr) - 1; end > 1; end-- {
		// sorted := true
		//
		sortedIndex := 0
		for j := 0; j < end; j++ {
			if arr[j] > arr[j+1] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
				// sorted = false
				sortedIndex = j + 1
			}
		}
		// if sorted {
		// 	break
		// }
		end = sortedIndex
	}
	fmt.Println(arr)
}
