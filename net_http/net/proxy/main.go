/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-13 22:31:56
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-14 22:16:43
 * @FilePath: /allfunc/net_http/net/proxy/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
)

type web1handler struct {
}

func (web1handler) ServeHTTP(writer http.ResponseWriter, r *http.Request) {

	writer.Write([]byte("web1"))
}

type web2handler struct {
}

func (web2handler) ServeHTTP(writer http.ResponseWriter, r *http.Request) {
	writer.Write([]byte("web2"))
}

func main() {
	c := make(chan os.Signal)

	defer func() {
		if err := recover(); err != nil {

			fmt.Println(err)
		}
	}()
	go func() {
		http.ListenAndServe(":9090", web1handler{})
	}()

	go func() {
		http.ListenAndServe(":9091", web2handler{})
	}()
	signal.Notify(c, os.Interrupt)
	s := <-c
	fmt.Println("777")
	fmt.Println(s)

}
