/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-10-31 16:16:11
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-10-31 16:16:54
 * @FilePath: /allfunc/gin_admin/common/mysql.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package common

import (
	"fmt"
	"log"
	"os"
	"project/allfunc/gin_admin/global"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

//创建mysql 连接
func NewMysqlConn() (db *gorm.DB) {
	db = GormMysql()
	return
}

func GormMysql() *gorm.DB {

	ms := global.GVA_CONFIG.MysqlConfig
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", ms.User, ms.Password, ms.Host, ms.Port, ms.Name)
	// root:%s@tcp(localhost:3306)/shop_db?charset=utf8mb4&parseTime=True&loc=Local"

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
