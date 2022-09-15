/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-14 22:48:13
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-14 22:51:17
 * @FilePath: /allfunc/base64/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AES
 */
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	str := "ykk 123"
	encode := base64.StdEncoding.EncodeToString([]byte(str))

	fmt.Println(encode)

	de, _ := base64.StdEncoding.DecodeString(encode)
	fmt.Println(string(de))
}
