/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-09 10:00:16
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-13 16:13:44
 * @FilePath: /allfunc/net_http/net/http/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE

 */
package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
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

	http.HandleFunc("/clientGet", curlStru.ClientGet)
	http.HandleFunc("/receiveClientGet", curlStru.ReceiveClientGet)
	http.HandleFunc("/clientPost", curlStru.ClientPost)
	http.HandleFunc("/receiveClientPost", curlStru.ReceiveClientPost)
	http.HandleFunc("/clientHouse", curlStru.LejuCurlInfo)
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

func (c *CurlStru) ClientGet(w http.ResponseWriter, r *http.Request) {

	data := r.URL.Query()

	name := data.Get("name")
	age := data.Get("age")
	client := &http.Client{}
	apiURL := "http://127.0.0.1:8090/receiveClientGet"

	req, err := http.NewRequest("GET", apiURL, nil)
	//添加查询参数
	q := req.URL.Query()
	q.Add("name", name)
	q.Add("age", age)
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	// type Cookie struct {
	// 	Name   Value  Path       Domain  Expires RawExpires MaxAge Secure HttpOnly SameSite Raw     Unparsed
	// }

	//设置cookie
	expire := time.Now().AddDate(0, 0, 1)

	cookie := http.Cookie{"test", "tcookie", "/", "www.domain.com", expire, expire.Format(time.UnixDate), 86400, true, true, 3, "test=tcookie", []string{"test=tcookie"}}
	req.Header.Add("Content-Type", "application/json")
	req.AddCookie(&cookie)
	if err != nil {
		fmt.Printf("post failed, err:%v\n\n", err)
		return
	}
	resp, _ := client.Do(req)
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed,err:%v\n\n", err)
		return
	}
	fmt.Println(string(b))
	w.Write(b)
}

func (c *CurlStru) ReceiveClientGet(w http.ResponseWriter, r *http.Request) {
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

func (c *CurlStru) ClientPost(w http.ResponseWriter, r *http.Request) {
	url := "http://127.0.0.1:8090/receiveClientPost"
	// 表单数据
	// contentType := "application/x-www-form-urlencoded"
	// data := "name=小王子&age=18"
	// json
	// contentType := "application/json"
	data := `{"name":"小王子","age":18}`
	cok, _ := r.Cookie("X-Xsrftoken")
	fmt.Println("cookie:", cok.Value)
	//方法一
	// resp, err := http.Post(url, contentType, strings.NewReader(data))

	//方法二
	// js, _ := json.Marshal(data)
	// js, _ := json.MarshalIndent(data, "", "   ") // 将json形式的字符串进行格式化

	req, _ := http.NewRequest("POST", url, strings.NewReader(data))

	// cookie2, _ := r.Cookie("name")
	cookie1 := &http.Cookie{Name: "X-Xsrftoken", Value: "df41ba54db5011e89861002324e63af81", HttpOnly: true, Secure: true}
	req.AddCookie(cookie1)
	// req.AddCookie(cookie2)
	// req, _ := http.NewRequest("POST", url, bytes.NewBuffer(js))
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Add()
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
	http.SetCookie(w, cookie1) //设置cookie数据
	fmt.Println(string(b))
	w.Write(b)
}

func (c *CurlStru) ReceiveClientPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	// 1. 请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.Method)
	fmt.Println(r.PostForm) // 打印form数据
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))
	cok, _ := r.Cookie("X-Xsrftoken")
	fmt.Println("cookie:", cok.Value)
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

type House struct {
	Type          string `json:"type"`
	Appid         string `json:"appid"`
	Pcount        string `json:"pcount"`
	Page          string `json:"page"`
	Count         string `json:"count"`
	Sign          string `json:"sign"`
	Relation_city string `json:"relation_city"`
	Status        string `json:"status"`
	Hid           string `json:"city"`
	Salestate     string `json:"salestate"`
}

func (c *CurlStru) LejuCurlInfo(w http.ResponseWriter, r *http.Request) {
	url := "http://test-info.leju.com/search/default/index"
	// 表单数据
	// contentType := "application/x-www-form-urlencoded"
	// data := "name=小王子&age=18"
	// json
	// contentType := "application/json"
	// data := `{"name":"小王子","age":18}`
	//方法一
	// resp, err := http.Post(url, contentType, strings.NewReader(data))

	var Parameters = map[string]interface{}{}
	//方法二
	// js, _ := json.Marshal(data)
	// js, _ := json.MarshalIndent(data, "", "   ") // 将json形式的字符串进行格式化
	// 	[type] => house
	// [appid] => 2016112598
	// [pcount] => 10
	// [page] => 1
	// [count] => 1
	// [filter1] => {status@eq}1
	// [filter2] => {hid@neq}0
	// [filter3] => {salestate@eq}1|2|3|10|4|11
	// [filter4] => {relation_city@eq}bj
	// [order] => {salestateorder@desc|house_hot_score@desc|updatetime@desc}desc
	// [sign] => 27aeeff73ce9e62b79ebd84adcce730d
	house := House{Type: "house", Appid: "2016112598", Page: "1", Pcount: "1", Count: "1", Relation_city: "ab", Hid: "132324", Salestate: "1|2|3|10|4|11", Status: "1"}
	appkey := "871196a770a984d5882c5a5df5a7494e"

	Parameters["type"] = house.Type
	Parameters["appid"] = house.Appid
	Parameters["pcount"] = house.Pcount
	Parameters["page"] = "1"
	Parameters["count"] = "1"
	Parameters["filter1"] = "{status@eq}1"
	Parameters["filter2"] = "{hid@eq}132324"
	Parameters["filter3"] = "{salestate@eq}1|2|3|10|4|11"
	Parameters["filter4"] = "{relation_city@eq}ab"
	Parameters["order"] = "{salestateorder@desc|house_hot_score@desc|updatetime@desc}desc"
	var keys []string = []string{"type", "appid", "pcount", "page", "count", "filter1", "filter2", "filter3", "filter4", "order", "sign"}

	fmt.Println("keys:", keys)
	str := ""
	for _, pa := range keys {
		if Parameters[pa] == nil {
			continue
		}
		str = str + fmt.Sprintf("%s", Parameters[pa])
	}
	fmt.Println("str:", str)
	m := md5.New()
	m.Write([]byte(str + appkey))
	md5str := hex.EncodeToString(m.Sum(nil))
	Parameters["sign"] = md5str
	st := HttpBuildQuery(Parameters, keys)

	req, _ := http.NewRequest("POST", url, strings.NewReader(st))
	fmt.Println("st:", st)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	// req.Header.Add()
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
	// http.SetCookie(w, cookie1) //设置cookie数据
	fmt.Println(string(b))
	w.Write(b)
}

func HttpBuildQuery(params map[string]interface{}, keys []string) (param_str string) {
	params_arr := make([]string, 0, len(params))
	for _, v := range keys {
		params_arr = append(params_arr, fmt.Sprintf("%s=%s", v, params[v]))
	}
	//fmt.Println(params_arr)
	param_str = strings.Join(params_arr, "&")
	return param_str
}
