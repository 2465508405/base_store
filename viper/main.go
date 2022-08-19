/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-05-22 22:00:57
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-19 17:24:24
 * @FilePath: /allfunc/viper/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
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

type MysqlJson struct {
	Port int    `mapstructure:"port"`
	User string `mapstructure:"user"`
}
type JsonConfig struct {
	Name      string `mapstructure:"name"`
	Age       string `mapstructure:"age"`
	MysqlJson `mapstructure:"mysql"`
}

func ReadConfigJson() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./")     // path to look for the config file in
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	sc := JsonConfig{}
	viper.Unmarshal(&sc)
	fmt.Println(sc)
	fmt.Println(viper.AllSettings())
	fmt.Println(viper.Get("name"))
}

func main() {
	// readConfig()
	// fmt.Println(GetEnvInfo("MXSHOP_DEBUG"))
	DevOnlineConfigDiv()

	// ReadConfigJson()
}
