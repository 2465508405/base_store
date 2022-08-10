/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-11 22:41:57
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-07-11 23:26:12
 *最大时间复杂度O(n^2)  平均 O(n^2) 最好时间复杂度 O(n)
 *空间复杂度 ：O(1)
 * 逆序对(左边数值大于右边数值的数量) ：数组<2,3,8,6,1>的逆序对为<2,1>, <3,1>,<8,6>,<8,1>,<6,1>
 * 插入排序的时间复杂度与逆序对的数量成正比
 */
package main

import "fmt"

func main() {

	var arr = []int{10, 2, 7, 9, 4, 11, 22322, 291, 23329, 131, 2, 29}

	for begin := 1; begin <= len(arr)-1; begin++ {
		val := arr[begin]
		currIndex := begin
		for end := begin - 1; end >= 0; end-- {
			end_val := arr[end]
			if end_val > val {
				arr[end+1] = end_val
				currIndex--
				// arr[end] = val
			} else {
				break
			}
		}
		arr[currIndex] = val
	}
	fmt.Println(arr)
}
