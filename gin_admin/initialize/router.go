/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-15 16:31:47
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-11-02 17:16:58
 * @FilePath: /allfunc/gin_admin/initialize/router.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package initialize

import (
	"project/allfunc/gin_admin/middlewares"
	"project/allfunc/gin_admin/routers"

	"github.com/gin-gonic/gin"
)

type Option func(*gin.RouterGroup)

var options = []Option{}

// 注册app的路由配置
// func Include(opts ...Option) {
// 	options = append(options, opts...)
// }

var WhiteRouter = []interface{}{"/login", "/auth/login", "/register", "/auth/register"}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("views/**/*")
	r.Static("/css", "public/css")
	r.Static("/js", "public/js")
	r.Static("/bootstrap", "public/bootstrap")
	r.Static("/jquery", "public/jquery")
	// r.Static("/public", http.Dir("./public"))
	// r.Use(middlewares.LoginMiddleWare())
	systemRouter := routers.RouterGroupApp.System

	// Include(systemRouter.OrderRouter.OrderRouter, UserRouter, InitHome, LoginRouter)

	// r.Use(middlewares.LoginMiddleWare(WhiteRouter))
	RouterGroup := r.Group("")
	RouterGroup.Use(middlewares.LoginMiddleWare(WhiteRouter))
	// RouterGroup.Use()
	{
		systemRouter.InitHomeRouter(RouterGroup)
		systemRouter.InitUserRouter(RouterGroup)
		systemRouter.InitLoginRouter(RouterGroup)
		systemRouter.InitOrderRouter(RouterGroup)
	}
	return r
}
