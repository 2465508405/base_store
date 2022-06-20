/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-06-19 21:31:13
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-06-19 21:31:23
 * @FilePath: /allfunc/oss/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func main() {

	fmt.Printf("oss-version:%s", oss.Version)

}
