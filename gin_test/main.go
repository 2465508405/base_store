/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-05-28 15:28:41
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-05-28 21:20:40
 * @FilePath: /gintest/redis_gin/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"project/allfunc/gin_test/routers"
	"runtime"
	"runtime/debug"

	"project/allfunc/gin_test/globals"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("rdb ", globals.RedisClient)
	// s, _ := globals.rdb.Set(context.Background(), "a", "33").Result()
	r := gin.Default()
	runtime.GOMAXPROCS(runtime.NumCPU()) //cpu使用
	debug.SetMaxThreads(30)              //设置线程数
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user", routers.GetUserList)
	r.Run(":8088") // 监听并在 0.0.0.0:8080 上启动服务
}
