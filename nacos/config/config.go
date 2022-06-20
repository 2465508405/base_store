/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-03-27 22:15:00
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-06-16 22:38:01
 * @FilePath: /mxshop-api/user-web/config/config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

type UserSrvConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
	Name string `mapstructure:"name" json:"name"`
}
type AliSmsConfig struct {
	ApiKey    string `mapstructure:"key" json:"key"`
	ApiSecret string `mapstructure:"secret" json:"secret"`
}
type RedisConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}
type JWTConfig struct {
	SigningKey string `mapstructure:"key" json:"key"`
}
type ConsulConfig struct {
	Host string `mapstructure:"host" json:"host"`
	Port int    `mapstructure:"port" json:"port"`
}
type ServerConfig struct {
	Name        string        `mapstructure:"name" json:"name"`
	Port        int           `mapstructure:"port" json:"port"`
	UserSrvInfo UserSrvConfig `mapstructure:"user_srv"  json:"user_srv"`
	JWTInfo     JWTConfig     `mapstructure:"jwt" json:"jwt"`
	AliSmsInfo  AliSmsConfig  `mapstructure:"sms" json:"sms"`
	RedisInfo   RedisConfig   `mapstructure:"redis" json:"redis"`
	ConsulInfo  ConsulConfig  `mapstructure:"consul" json:"consul"`
}
