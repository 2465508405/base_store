/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-14 22:06:55
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-22 10:25:55
 * @FilePath: /allfunc/net_http/net/proxy/proxy/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"project/allfunc/net_http/net/proxy/util"
)

type Proxy struct {
}

func (Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if err := recover(); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Println(err)
		}
	}()
	url := r.URL.Path
	fmt.Println(url)
	if url == "/a" {
		fmt.Println(r.Method)
		fmt.Println(r.RemoteAddr)
		// auth := r.Header.Get("Authorization")
		newreq, _ := http.NewRequest(r.Method, "http://localhost:9090", r.Body)
		// newreq.Header.Set("Authorization", auth)
		fmt.Println(r.Header, newreq.Header)
		util.CloneHeader(r.Header, &newreq.Header)
		newreq.Header.Add("x-forwarded-for", r.RemoteAddr) //
		fmt.Println(newreq.Header)
		response, _ := http.DefaultClient.Do(newreq)
		fmt.Println(response)
		getHeader := w.Header()
		w.Header().Set("Content-Type", "application/json;charset=utf-8")
		fmt.Println(w.Header())
		util.CloneHeader(response.Header, &getHeader) //拷贝头部

		w.WriteHeader(response.StatusCode) //写入http状态
		defer response.Body.Close()
		rsp, _ := io.ReadAll(response.Body)

		w.Write(rsp)
		return
	}
	w.Write([]byte("default index"))
}

func main() {

	ch := make(chan os.Signal)

	defer func() {
		if err := recover(); err != nil {

			fmt.Println(err)
		}
	}()

	go func() {
		http.ListenAndServe(":8080", Proxy{})
	}()
	signal.Notify(ch, os.Interrupt)
	s := <-ch
	fmt.Println("close:s->", s)
}
