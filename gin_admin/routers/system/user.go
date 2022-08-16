/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-17 12:44:17
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-15 17:41:48
 * @FilePath: /allfunc/leju_test/routers/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

import (
	v1 "project/allfunc/gin_admin/api/v1"
	"project/allfunc/gin_admin/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ua *UserRouter) InitUserRouter(Router *gin.RouterGroup) {

	routerGroup := Router.Group("user").Use(middlewares.OperationRecord())
	userApi := v1.ApiGroupApp.SystemApiGroup.UserApi
	{

		routerGroup.GET("/add", userApi.UserAdd)
		routerGroup.GET("/edit", userApi.UserEdit)
		routerGroup.GET("/list", userApi.UserList)
		routerGroup.POST("/postCreate", userApi.UserCreate)
		routerGroup.POST("/postEdit", userApi.UserUpdate)
		routerGroup.POST("/postDel", userApi.UserDel)
		// routerGroup.POST("/file/upload", api.FileUpload)
	}

}
