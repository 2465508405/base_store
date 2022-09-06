/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-02 16:05:24
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-02 17:37:00
 * @FilePath: /allfunc/sort/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"sort"
)

type sint []int32

func main() {

	var s sint
	s = sint{2, 3, 3, 3, 43, 23, 2332, 23, 23, 23, 22, 23, 23, 23, 23, 23, 23, 23, 1, 1}
	fmt.Println("cap:", cap(s))
	sortedHashes := units{1, 2, 34, 2923, 2923, 3232, 222, 323, 323, 23, 2, 1, 5, 5525, 41, 66, 626, 36, 36}

	//查找字符串
	str := []string{"a", "b", "c"}

	ind := sort.SearchStrings(str, "c")
	fmt.Println("str index:", ind)
	sort.Sort(sortedHashes) //重写排序方法，Len, Less, Swap

	fmt.Println(sortedHashes)
	f := func(i int) bool {

		return sortedHashes[i] > 10
	}

	in := sort.Search(len(sortedHashes), f)
	fmt.Println(in)
}

type units []int32

//返回切片长度
func (x units) Len() int {
	return len(x)
}

//比对两个数大小
func (x units) Less(i, j int) bool {
	return x[i] < x[j]
}

//切片中两个值的交换
func (x units) Swap(i, j int) {
	x[i], x[j] = x[j], x[i]
}
