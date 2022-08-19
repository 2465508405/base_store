/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-16 22:20:41
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-16 22:34:33
 * @FilePath: /allfunc/gin_admin/config/system.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package config

type System struct {
	DbType string `mapstructure:"db-type" json:"db-type" yaml:"db-type"` // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
}
