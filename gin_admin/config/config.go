/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-16 22:17:39
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-22 18:15:32
 * @FilePath: /allfunc/gin_admin/config/config.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

type ServerConfig struct {
	System      `mapstructure:"system" json:"system" yaml:"system"` // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	MysqlConfig `mapstructure:"mysql" json:"mysql" yaml:"mysql"`    // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
}
