/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-13 22:31:56
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-22 10:24:09
 * @FilePath: /allfunc/net_http/net/proxy/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
)

type web1handler struct {
}

func (wh web1handler) GetIP(r *http.Request) string {
	ips := r.Header.Get("x-forwarded-for")
	if ips != "" {
		ips_list := strings.Split(ips, ",")
		if len(ips_list) > 0 && ips_list[0] != "" {
			return ips_list[0]
		}
	}

	return r.RemoteAddr
}

func (wh web1handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	auth := r.Header.Get("Authorization")
	if auth == "" {
		w.Header().Set("WwW-Authenticate", `Basic realm="您必须输入用户名和密码"`)
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	fmt.Println(auth)
	auth_list := strings.Split(auth, " ")
	fmt.Println(auth_list)
	if len(auth_list) == 2 && auth_list[0] == "Basic" {
		info, _ := base64.StdEncoding.DecodeString(auth_list[1])
		fmt.Println(string(info))
		if string(info) == "ykk:123" {
			w.Write([]byte(fmt.Sprintf("<h2>web来自于：%s</h2>", wh.GetIP(r))))
			// w.Write([]byte(fmt.Sprintf("<h2>web来自于：%s</h2>", r.RemoteAddr)))
			return
		}
	}
	w.Write([]byte("用户名密码有误"))
	// writer.Write([]byte("web1"))
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
