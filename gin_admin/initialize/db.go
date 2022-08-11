package initialize

import (
	"log"
	"os"
	"project/allfunc/gin_admin/global"
	"project/allfunc/gin_admin/models"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func InitDB() {

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
	global.DB = db

	// 自动迁移
	db.AutoMigrate(&models.User{})
}
