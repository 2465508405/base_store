/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-16 22:41:48
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-17 22:28:36
 * @FilePath: /allfunc/gin_admin/config/mysql.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

type MysqlConfig struct {
	Host     string `mapstructure:"host" json:"host"`
	Port     int    `mapstructure:"port" json:"port"`
	Name     string `mapstructure:"db" json:"db"`
	User     string `mapstructure:"user" json:"user"`
	Password string `mapstructure:"password" json:"password"`
}
