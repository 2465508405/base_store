/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-10 17:10:43
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-19 17:01:43
 * @FilePath: /allfunc/gin_admin/global/global.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package global

import (
	"project/allfunc/gin_admin/config"

	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_CONFIG *config.ServerConfig = &config.ServerConfig{}
	// ServerConfig config.ServerConfig
	// NacosConfig  *config.NacosConfig
)
