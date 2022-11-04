/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-11-04 10:46:49
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-11-04 10:47:01
 * @FilePath: /allfunc/gin_admin/config/redis.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

type RedisConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"name" json:"name"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}
