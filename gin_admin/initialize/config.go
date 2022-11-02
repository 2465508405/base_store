/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-17 22:03:58
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-10-28 14:11:41
 * @FilePath: /allfunc/gin_admin/initialize/config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package initialize

import (
	"encoding/json"
	"fmt"
	"project/allfunc/gin_admin/global"

	"github.com/fsnotify/fsnotify"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func GetEnvInfo(s string) bool {
	viper.AutomaticEnv()

	return viper.GetBool(s)
}

func InitConfig() {

	env := GetEnvInfo("MXSHOP_DEBUG")

	prefix := "config"
	fileName := fmt.Sprintf("../gin_admin/%s.yaml", prefix)
	if env {
		fileName = fmt.Sprintf("../gin_admin/%s-dev.yaml", prefix)
	}
	fmt.Println(fileName)
	vp := viper.New()

	vp.SetConfigFile(fileName)
	if err := vp.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired

		} else {
			// Config file was found but another error was produced
		}
		fmt.Println(err)
	}
	// fmt.Println(vp.AllSettings())
	// var serverConfig = config.ServerConfig{}

	vp.Unmarshal(&global.GVA_NACOS)
	// fmt.Println(global.GVA_NACOS)
	InitGlobalConfig()
	//动态监控变化
	vp.WatchConfig()
	vp.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件产生变化:%s", e.Name)
		fmt.Println("config file changed", e.Name)
		_ = vp.ReadInConfig()
		_ = vp.Unmarshal(&global.GVA_NACOS)
		InitGlobalConfig()
		zap.S().Infof("配置信息:&v", global.GVA_NACOS)
	})
	// time.Sleep(time.Second * 30)

}

//从Nacos中获取数据
func InitGlobalConfig() {
	sc := []constant.ServerConfig{
		{
			IpAddr: global.GVA_NACOS.Host,
			// ContextPath: "/nacos",
			Port: uint64(global.GVA_NACOS.Port),
			// Scheme:      "http",
		},
	}

	// 创建clientConfig
	cc := constant.ClientConfig{
		NamespaceId:         "5e1c2dda-e6fb-4272-bc36-7790f9f04f32", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	if err != nil {
		panic(err)
	}

	content, _ := configClient.GetConfig(vo.ConfigParam{
		DataId: "shop_db.json",
		Group:  "dev",
	})

	json.Unmarshal([]byte(content), &global.GVA_CONFIG)
	fmt.Printf("gva_config :%+v\n", global.GVA_CONFIG)
}
