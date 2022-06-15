/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-06-01 16:18:53
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-06-06 16:01:31
 * @FilePath: /allfunc/elastic/op/elastic.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package op

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
)

type Article struct {
	Title   string    // 文章标题
	Content string    // 文章内容
	Author  string    // 作者
	Created time.Time // 发布时间
}

var es *elasticsearch.Client

func init() {

	cfg := elasticsearch.Config{
		Addresses: []string{
			"http://localhost:9200",
		},
		Username: "elastic",
		Password: "123456",
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS11,
			},
		},
	}
	es, _ = elasticsearch.NewClient(cfg)
}
func Conn() {

	// 测试
	_, err := es.Ping()
	if err != nil {
		panic(err)
	}

	// 打印es版本信息
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	log.Println(res)
}
