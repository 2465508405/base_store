/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-17 22:03:58
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-19 17:29:58
 * @FilePath: /allfunc/gin_admin/initialize/config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package initialize

import (
	"fmt"
	"project/allfunc/gin_admin/global"

	"github.com/fsnotify/fsnotify"
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

	vp.Unmarshal(&global.GVA_CONFIG)
	fmt.Println(global.GVA_CONFIG)
	//动态监控变化
	vp.WatchConfig()
	vp.OnConfigChange(func(e fsnotify.Event) {
		zap.S().Infof("配置文件产生变化:%s", e.Name)
		fmt.Println("config file changed", e.Name)
		_ = vp.ReadInConfig()
		_ = vp.Unmarshal(&global.GVA_CONFIG)
		zap.S().Infof("配置信息:&v", global.GVA_CONFIG)
	})
	// time.Sleep(time.Second * 30)

}
