/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-05-21 20:39:30
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-05-21 21:01:29
 * @FilePath: /allfunc/flag/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"flag"
	"fmt"
)

var flagvar int

func main() {

	flag.IntVar(&flagvar, "flagname", 1234, "help message for flagname") //cmd  go run main.go -flagname=22322
	ip := flag.String("ip", "0.0.0.0", "ip地址")                           // go run main.go -ip=223.2.21.2
	flag.Parse()                                                         //解析输入的命令行参数
	fmt.Println(*ip)
	fmt.Println(flagvar)
}
