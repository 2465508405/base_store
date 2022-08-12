/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-17 12:44:17
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-12 10:14:50
 * @FilePath: /allfunc/leju_test/routers/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package routers

import (
	"project/allfunc/gin_admin/api"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) {

	r.GET("/user/list", api.UserList)
	r.GET("/user/add", api.UserAdd)
	r.GET("/user/edit", api.UserEdit)

	r.POST("/user/postCreate", api.UserCreate)
	r.POST("/user/postEdit", api.UserUpdate)
	r.POST("/user/postDel", api.UserDel)
	r.POST("/file/upload", api.FileUpload)

}
