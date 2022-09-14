/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-02 14:18:44
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-13 16:54:22
 * @FilePath: /allfunc/protobuf_test/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"encoding/json"
	"fmt"
	"project/allfunc/protobuf_test/proto"

	pt "google.golang.org/protobuf/proto"
)

type Hello struct {
	Name    string   `json:"name"`
	Age     int32    `json:"age"`
	Courses []string `json:"courses"`
}

func main() {

	tt := proto.HelloRequest{}
	fmt.Printf("%s", tt.Name)
	fmt.Println("========")
	req := proto.HelloRequest{
		Name:    "bobby",
		Age:     18,
		Courses: []string{"go", "gin", "微服务"},
	}
	jsonStruct := Hello{
		Name:    "bobby",
		Age:     18,
		Courses: []string{"go", "gin", "微服务"},
	}

	//比 json数据压缩比更高
	jsonRsp, _ := json.Marshal(jsonStruct)
	fmt.Println(len(jsonRsp))
	rsp, _ := pt.Marshal(&req)
	fmt.Println(len(rsp))
	newReq := proto.HelloRequest{}
	_ = pt.Unmarshal(rsp, &newReq)
	fmt.Printf("name:%s, age:%d, courses:%v", newReq.Name, newReq.Age, newReq.Courses)

}
