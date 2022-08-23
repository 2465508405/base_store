/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-06-15 22:56:12
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-22 18:07:14
 * @FilePath: /allfunc/nacos/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"project/allfunc/nacos/config"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type str struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}

func main() {
	// NacosOne()
	// NacosTwo()
	NacosThree()
}

//服务注册，创建，删除
func NacosTwo() {

	sc := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			// ContextPath: "/nacos",
			Port: 8848,
			// Scheme:      "http",
		},
	}

	// 创建clientConfig
	cc := constant.ClientConfig{
		NamespaceId:         "007214bb-b65f-460b-bf97-86f751caf94f", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建服务发现客户端的另一种方式 (推荐)
	namingClient, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}

	//注册服务实例
	// success, err := namingClient.RegisterInstance(vo.RegisterInstanceParam{
	// 	Ip:          "127.0.0.1",
	// 	Port:        8848,
	// 	ServiceName: "demo.go",
	// 	Weight:      10,
	// 	Enable:      true,
	// 	Healthy:     true,
	// 	Ephemeral:   true,
	// 	Metadata:    map[string]string{"idc": "shanghai"},
	// 	ClusterName: "cluster-a", // 默认值DEFAULT
	// 	GroupName:   "group-a",   // 默认值DEFAULT_GROUP
	// })
	//注销实例：DeregisterInstance
	// success, err := namingClient.DeregisterInstance(vo.DeregisterInstanceParam{
	// 	Ip:          "127.0.0.1",
	// 	Port:        8848,
	// 	ServiceName: "demo.go",
	// 	Ephemeral:   true,
	// 	Cluster:     "cluster-a", // 默认值DEFAULT
	// 	GroupName:   "group-a",   // 默认值DEFAULT_GROUP
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// if success {
	// 	fmt.Println("success register")
	// }
	//获取所有的实例列表：SelectAllInstances
	// instances, err := namingClient.SelectAllInstances(vo.SelectAllInstancesParam{
	// 	ServiceName: "demo.go",
	// 	GroupName:   "group-a",             // 默认值DEFAULT_GROUP
	// 	Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
	// })
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(333)
	// for _, inst := range instances {
	// 	fmt.Println(inst)
	// }
	// fmt.Println(77777)

	//监听服务变化：Subscribe
	// Subscribe key=serviceName+groupName+cluster
	// 注意:我们可以在相同的key添加多个SubscribeCallback.
	err = namingClient.Subscribe(&vo.SubscribeParam{
		ServiceName: "demo.go",
		GroupName:   "group-a",             // 默认值DEFAULT_GROUP
		Clusters:    []string{"cluster-a"}, // 默认值DEFAULT
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			// log.Printf("\n\n callback return services:%s \n\n", utils.ToJsonString(services))
			log.Printf("\n\n callback return services:%v \n\n", services)
		},
	})
	if err != nil {
		panic(err)
	}

	time.Sleep(3000 * time.Second)
}

//创建动态配置客户端
func NacosThree() {

	sc := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			// ContextPath: "/nacos",
			Port: 8848,
			// Scheme:      "http",
		},
	}

	// 创建clientConfig
	cc := constant.ClientConfig{
		NamespaceId:         "007214bb-b65f-460b-bf97-86f751caf94f", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建动态配置客户端
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}

	s := str{Host: "127.0.0.1", Port: 9222}
	sj, err := json.Marshal(s)
	//发布配置：PublishConfig
	success, err := configClient.PublishConfig(vo.ConfigParam{
		DataId:  "dataId",
		Group:   "group",
		Content: string(sj),
	})

	//删除配置：DeleteConfig

	// success, err := configClient.DeleteConfig(vo.ConfigParam{
	// 	DataId: "dataId",
	// 	Group:  "group"})

	if success {
		fmt.Println("success delete")
	}
	//获取配置
	// content, err := configClient.GetConfig(vo.ConfigParam{
	// 	DataId: "user-web.json",
	// 	Group:  "dev"})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(content)
	// serverConfig := config.ServerConfig{}
	// json.Unmarshal([]byte(content), &serverConfig)
	// fmt.Printf("%+v\n", serverConfig)

	// err = configClient.ListenConfig(vo.ConfigParam{
	// 	DataId: "user-web.json",
	// 	Group:  "dev",
	// 	OnChange: func(namespace, group, dataId, data string) {
	// 		fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
	// 	},
	// })
	// if err != nil {
	// 	panic(err)
	// }
	time.Sleep(3000 * time.Second)
}

func NacosOne() {

	sc := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			// ContextPath: "/nacos",
			Port: 8848,
			// Scheme:      "http",
		},
	}

	// 创建clientConfig
	cc := constant.ClientConfig{
		NamespaceId:         "007214bb-b65f-460b-bf97-86f751caf94f", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	// 创建动态配置客户端
	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "dev"})
	if err != nil {
		panic(err)
	}
	// fmt.Println(content)
	serverConfig := config.ServerConfig{}
	json.Unmarshal([]byte(content), &serverConfig)
	fmt.Println(serverConfig)

	err = configClient.ListenConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "dev",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
	if err != nil {
		panic(err)
	}
	time.Sleep(3000 * time.Second)
}
