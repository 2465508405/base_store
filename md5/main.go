package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
)

/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-22 13:50:38
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-11 15:11:41
 * 加密数据与  php 保持一致
 */

func main() {

	//方法1
	md1()

	md2()

	md3()

	md4()
}

func md1() {
	//方法1
	str := "123456abc777"
	m := md5.New()
	m.Write([]byte(str))
	fmt.Println(hex.EncodeToString(m.Sum(nil)))
}

func md2() {
	//方法2
	str := "123456abc777"
	m := md5.Sum([]byte(str))
	fmt.Println(hex.EncodeToString(m[:]))
}

func md3() {
	str := "123456abc777"
	m := md5.Sum([]byte(str))
	fmt.Printf("%x\n", m)

}

func md4() {
	str := "123456abc777"
	m := md5.New()
	io.WriteString(m, str)
	fmt.Println(hex.EncodeToString(m.Sum(nil)))
}
