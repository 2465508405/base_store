/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-10 18:46:00
 * @LastEditors: ykk ykk@qq.com
*
*从序列中找到最大的元素，然后与最末尾的元素进行交换
*
*/
package main

import "fmt"

func main() {

	arr := []int{2, 322, 232, 231, 92, 323, 9292, 2332, 931}

	for end := len(arr) - 1; end > 1; end-- {
		maxIndex := end
		for start := 0; start < end-1; start++ {
			if arr[start] > arr[maxIndex] {
				maxIndex = start
			}
		}
		if maxIndex != end {
			arr[end], arr[maxIndex] = arr[maxIndex], arr[end]
		}
	}
	fmt.Println(arr)
}
