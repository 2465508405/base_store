/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-17 20:07:02
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-07-19 23:00:48
 * @FilePath: /allfunc/leju_test/lib/curl/http.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package curl

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

func HTTPJson(method string, url string, params []byte) ([]byte, error) {
	body := bytes.NewBuffer(params)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp != nil {
		defer resp.Body.Close()
	}
	if resp.StatusCode != http.StatusOK {
		msg := fmt.Sprintf("response status code %v", resp.StatusCode)
		return nil, errors.New(msg)
	}

	bs, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(bs))
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func HTTPDo(method string, url string, values url.Values) ([]byte, error) {
	body := strings.NewReader(values.Encode())

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	//req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	//req.Header.Set("Cookie", cookie)
	//req.Header.Set("Connection", "keep-alive")
	//req.Header.Add("x-requested-with", "XMLHttpRequest") //AJAX

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	if resp.StatusCode == http.StatusOK {
		msg := fmt.Sprintf("response status code %v", resp.StatusCode)
		return nil, errors.New(msg)
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func HTTPGet(url string) ([]byte, error) {
	//发送请求获取响应
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//结束网络释放资源
	if resp != nil {
		defer resp.Body.Close()
	}
	//判断响应状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(fmt.Sprintf("response status code %v", resp.StatusCode))
	}

	//读取响应实体
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func HTTPPost(url string, params url.Values, contentType string) ([]byte, error) {

	body := strings.NewReader(params.Encode())

	if contentType == "" {
		contentType = "application/x-www-form-urlencoded"
	}

	resp, err := http.Post(url, contentType, body)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(fmt.Sprintf("response status code %v", resp.StatusCode))
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func HTTPPostForm(url string, values url.Values) ([]byte, error) {
	resp, err := http.PostForm(url, values)
	if err != nil {
		return nil, err
	}
	if resp != nil {
		defer resp.Body.Close()
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(fmt.Sprintf("response status code %v", resp.StatusCode))
	}

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func MakeParams(params url.Values) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	bb := bytes.Buffer{}
	for _, key := range keys {
		val := params.Get(key)
		bb.WriteString(key)
		bb.WriteString("=")
		bb.WriteString(val)
		bb.WriteString("&")
	}

	return strings.TrimRight(bb.String(), "&")
}
