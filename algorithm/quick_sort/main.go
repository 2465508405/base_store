/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-13 23:00:54
 * @LastEditors: ykk ykk@qq.com
 *平均时间复杂度为O（nlogn）, 不稳定排序
 */
package main

import "fmt"

func main() {
	arr := []int{2323, 23232, 21, 6, 2, 5, 545, 84, 34, 2, 8, 34, 37}
	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(nums []int, l, r int) { //[l,r]
	if l < r {
		m := partition(nums, l, r)
		quickSort(nums, l, m-1)
		quickSort(nums, m+1, r)
	}
}

func partition(nums []int, l int, r int) int {
	key := nums[r]
	//all in [l,i) < key
	//all in [i,j] > key
	i := l
	j := l
	for j < r {
		if nums[j] < key {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}
	nums[i], nums[r] = nums[r], nums[i]
	return i
}
