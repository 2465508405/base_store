/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-05-30 14:49:14
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2023-02-23 17:03:55
 * @FilePath: /allfunc/consul/consul_test/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package consul_test

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

// 获取所有服务
func AllService() {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}

	for key, _ := range data {
		fmt.Println("key:", key)
	}
}

// 过滤服务
func FilterService() {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	data, err := client.Agent().ServicesWithFilter(`Service == "user-web"`)
	if err != nil {
		panic(err)
	}

	for key, _ := range data {
		fmt.Println("key:", key)
	}
}
func Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		HTTP:                           "http://www.baidu.com",
		Interval:                       "3s", //必要参数
		Timeout:                        "5s",
		DeregisterCriticalServiceAfter: "10s",
	}
	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.Port = port
	registration.Address = address
	registration.Tags = tags
	registration.ID = id
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
	fmt.Println("register success")
	return nil
}
