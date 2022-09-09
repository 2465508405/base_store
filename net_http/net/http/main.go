/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-09 10:00:16
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-09 14:55:52
 * @FilePath: /allfunc/net_http/net/http/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE

 */
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type CurlStru struct {
}

type User struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

type SuccessInfo struct {
	Status string `json:"status"`
	User   User   `json:"data"`
}

var curlStru = CurlStru{}

func main() {

	http.HandleFunc("/httpGet", curlStru.HttpGet)
	http.HandleFunc("/receiveHttpGet", curlStru.ReceiveHttpGet)
	http.HandleFunc("/httpPost", curlStru.HttpPost)
	http.HandleFunc("/receiveHttpPost", curlStru.ReceiveHttpPost)
	http.ListenAndServe(":8090", nil)

}

func (c *CurlStru) HttpGet(w http.ResponseWriter, r *http.Request) {

	apiUrl := "http://localhost:8090/receiveHttpGet"

	// URL param 地址传参数
	data := url.Values{}
	data.Set("name", "小王子")
	data.Set("age", "18")
	u, err := url.ParseRequestURI(apiUrl)

	if err != nil {
		fmt.Printf("parse url requestUrl failed,err:%v\n", err)
	}
	u.RawQuery = data.Encode() // URL encode Encode encodes the values into “URL encoded” form ("bar=baz&foo=quux") sorted by key.
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		panic(err)
	}
	succ := new(SuccessInfo)

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	json.Unmarshal(body, succ)
	fmt.Printf("%+v", succ)
	w.Write(body)
}

func (c *CurlStru) ReceiveHttpGet(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	user := User{Name: data.Get("name"), Age: data.Get("age")}

	success := SuccessInfo{Status: "ok", User: user}
	success_info, _ := json.Marshal(success)

	// answer := `{"status": "ok", "data":{"name":"` + data.Get("name") + `}}`
	w.Write(success_info)
}

func (c *CurlStru) HttpPost(w http.ResponseWriter, r *http.Request) {
	url := "http://127.0.0.1:8090/receiveHttpPost"
	// 表单数据
	// contentType := "application/x-www-form-urlencoded"
	// data := "name=小王子&age=18"
	// json
	// contentType := "application/json"
	data := `{"name":"小王子","age":18}`
	//方法一
	// resp, err := http.Post(url, contentType, strings.NewReader(data))

	//方法二
	// js, _ := json.Marshal(data)
	// js, _ := json.MarshalIndent(data, "", "   ") // 将json形式的字符串进行格式化
	req, _ := http.NewRequest("POST", url, strings.NewReader(data))
	// req, _ := http.NewRequest("POST", url, bytes.NewBuffer(js))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("get resp failed,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	w.Write(b)
}

func (c *CurlStru) ReceiveHttpPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.Method)
	fmt.Println(r.PostForm) // 打印form数据
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	// 2. 请求类型是application/json时从r.Body读取数据
	b, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	answer := `{"status": "ok"}`
	w.Write([]byte(answer))
}
