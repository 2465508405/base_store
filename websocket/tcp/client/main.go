/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-26 16:44:48
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-26 19:03:38
 * @FilePath: /allfunc/websocket/tcp/client/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"flag"
	"fmt"
	"net"
	"time"
)

var flagvar string

func main() {
	flag.CommandLine.StringVar(&flagvar, "flagname", "", "help message for flagname")
	flag.Parse()
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you? ---` + flagvar
		conn.Write([]byte(msg))
		time.Sleep(time.Second * 2)
	}
}
