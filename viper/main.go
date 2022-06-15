package main

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}
type ServerConfig struct {
	ServiceName string      `mapstructure:"name"`
	MysqlConfig MysqlConfig `mapstructure:"mysql"`
}

//读取配置文件内容
func readConfig() {
	v := viper.New()
	v.SetConfigFile("./config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	// v.Get('name')//获取配置文件中的变量值
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig.MysqlConfig)
}

func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

//配置文件隔离
func DevOnlineConfigDiv() {

	debug := GetEnvInfo("MXSHOP_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("%s.yaml", configFilePrefix)
	if debug {
		configFileName = fmt.Sprintf("%s-dev.yaml", configFilePrefix)
	}

	v := viper.New()
	v.SetConfigFile(configFileName)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//fmt.Println(v.Get("name")) //获取配置文件中的变量值
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil {
		panic(err)
	}
	fmt.Println(serverConfig.MysqlConfig)

	//监听配置文件变化
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&serverConfig)
		fmt.Println(serverConfig)
	})
	v.WatchConfig()

	time.Sleep(30 * time.Second)

}
func main() {
	// readConfig()
	// fmt.Println(GetEnvInfo("MXSHOP_DEBUG"))
	DevOnlineConfigDiv()
}
