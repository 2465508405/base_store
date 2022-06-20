/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-06-15 22:56:12
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-06-16 22:41:22
 * @FilePath: /allfunc/nacos/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"encoding/json"
	"fmt"
	"project/allfunc/nacos/config"
	"time"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {

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
