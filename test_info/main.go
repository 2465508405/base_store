/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-15 14:16:49
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-17 22:56:30
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

func main() {

	pa := new(Parent)

	fmt.Println(pa)
}
