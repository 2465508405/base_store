/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-10 10:59:34
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-16 17:37:52
 * @FilePath: /allfunc/gin_admin/routers/home.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

import (
	v1 "project/allfunc/gin_admin/api/v1"

	"github.com/gin-gonic/gin"
)

type HomeRouter struct{}

func (ha *HomeRouter) InitHomeRouter(r *gin.RouterGroup) {

	homeApi := v1.ApiGroupApp.SystemApiGroup.HomeApi
	r.GET("/home", homeApi.Home)
}
