/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-05-30 14:49:14
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-06-09 16:28:39
 * @FilePath: /allfunc/consul/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"project/allfunc/consul/consul_test"
)

func main() {
	consul_test.Register("172.17.0.2", 8021, "user-web", []string{"mxshop", "bobby"}, "user-web-1")
	// consul_test.AllService()
	// consul_test.AllService()
	consul_test.FilterService()
	// svrID := consul_test.HttpReg("mysvr-rpc", "127.0.0.1", 3456)
	// log.Printf("服务注册成功：%v\n", svrID)
}
