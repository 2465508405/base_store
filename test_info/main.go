/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-15 14:16:49
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-02 14:54:12
 * @FilePath: /test_info/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
)

type Parent struct {
	Stu
}
type Stu struct {
	Id  int
	Age string
}

func (s *Stu) add() {
	fmt.Println("aasss")
}

var h int

var mp []int

func main() {

	pa := new(Parent)

	fmt.Println(pa)

	mp = make([]int, 0)
	mp = append(mp, 2)

	fmt.Println(mp)
	key := "abc阿啊发发"
	var srcatch []byte
	//拷贝数据到数组中
	/**
	*copy时，必须指定切片长度
	 */
	srcatch = make([]byte, len(key))
	copy(srcatch, key)
	fmt.Println(srcatch)
	fmt.Println([]byte(key))

	/* 拷贝 numbers 的内容到 numbers1 */
	numbers := []int{5, 6, 7}
	numbers1 := make([]int, len(numbers)+3, (cap(numbers))*2)
	copy(numbers1, numbers)
	fmt.Println(numbers1)
}
