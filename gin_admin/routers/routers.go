/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-17 12:21:00
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-12 10:15:49
 * @FilePath: /allfunc/leju_test/router/router.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package routers

import (
	"github.com/gin-gonic/gin"
)

type Option func(*gin.Engine)

var options = []Option{}

// 注册app的路由配置
func Include(opts ...Option) {
	options = append(options, opts...)
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("views/**/*")
	r.Static("/css", "public/css")
	r.Static("/js", "public/js")
	r.Static("/bootstrap", "public/bootstrap")
	r.Static("/jquery", "public/jquery")
	// r.Static("/public", http.Dir("./public"))
	// r.Use(middlewares.LoginMiddleWare())
	Include(UserRouter, InitHome, LoginRouter)
	for _, opt := range options {
		opt(r)
	}
	return r
}
