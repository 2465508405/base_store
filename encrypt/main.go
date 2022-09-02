package main

import (
	"fmt"
	"project/allfunc/encrypt/encrypt"
)

func main() {

	str := "1235abc哈哈"
	enstr, _ := encrypt.EnPwdCode([]byte(str))

	fmt.Println(enstr)
	strByte, _ := encrypt.DePwdCode(enstr)

	fmt.Println(string(strByte))
}
