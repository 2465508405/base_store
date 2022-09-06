/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-04 09:41:54
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-04 10:01:59
 * @FilePath: /allfunc/net_http/net/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	resp, err := http.Get("http://www.baidu.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	fmt.Println(string(body))

	client := http.Client{}
	resp, err = client.Get("http://www.baidu.com")
	http.HandleFunc("/haha", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("nihao"))
	})
	log.Fatalln(http.ListenAndServe(":8090", nil))
}
