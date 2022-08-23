/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-22 17:42:56
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-22 17:44:24
 * @FilePath: /allfunc/gin_admin/config/nacos.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

type NacosConfig struct {
	Host string `mapstructure:"host" json:"host" yaml:"host"`
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
}
