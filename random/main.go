/*
 * @Author: ykk ykk@qq.com
 * @Date: 2023-02-23 15:09:01
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2023-02-23 15:09:10
 * @FilePath: /allfunc/random/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"math/rand"
	"strings"
)

var CHARS = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
	"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z",
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}

/*
RandAllString  生成随机字符串([a~zA~Z0~9])

	lenNum 长度
*/
func RandAllString(lenNum int) string {
	str := strings.Builder{}
	length := len(CHARS)
	for i := 0; i < lenNum; i++ {
		l := CHARS[rand.Intn(length)]
		str.WriteString(l)
	}
	return str.String()
}

/*
RandNumString  生成随机数字字符串([0~9])

	lenNum 长度
*/
func RandNumString(lenNum int) string {
	str := strings.Builder{}
	length := 10
	for i := 0; i < lenNum; i++ {
		str.WriteString(CHARS[52+rand.Intn(length)])
	}
	return str.String()
}

/*
RandString  生成随机字符串(a~zA~Z])

	lenNum 长度
*/
func RandString(lenNum int) string {
	str := strings.Builder{}
	length := 52
	for i := 0; i < lenNum; i++ {
		str.WriteString(CHARS[rand.Intn(length)])
	}
	return str.String()
}

func main() {
	s := RandString(6)
	fmt.Println(s)
}
