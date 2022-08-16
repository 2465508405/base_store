/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-10 22:28:34
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-16 17:19:48
 * @FilePath: /gin_admin/initialize/db.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package initialize

import (
	"fmt"
	"log"
	"os"
	"time"

	"project/allfunc/gin_admin/models/system"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() *gorm.DB {

	dsn := "root:@tcp(localhost:3306)/shop_db?charset=utf8mb4&parseTime=True&loc=Local"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢sql
			LogLevel:      logger.Info, //日志级别
			Colorful:      true,        //禁用彩色打印
		},
	)

	//全局模式 打印日志
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //根据结构体名声称表，表名不加s
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	return db
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
