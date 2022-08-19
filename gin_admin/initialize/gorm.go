/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-10 22:28:34
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-16 22:16:03
 * @FilePath: /gin_admin/initialize/db.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package initialize

import (
	"fmt"
	"os"

	"project/allfunc/gin_admin/global"
	"project/allfunc/gin_admin/models/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	switch global.GVA_CONFIG.System.DbType {
	case "mysql":
		return GormMysql()
	// case "pgsql":
	// 	return GormPgSql()
	default:
		return GormMysql()
	}
}

//自动迁移
func MigrateTables(db *gorm.DB) {
	err := db.AutoMigrate(
		// 系统模块表
		system.User{},
		system.Banner{},
		system.Goods{},
		system.Orders{},
	)
	if err != nil {
		fmt.Printf("register table failed: %v\n", zap.Error(err))
		os.Exit(0)
	}
	fmt.Println("register table success")
}
