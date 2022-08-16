/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-13 16:35:00
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-15 17:41:14
 * @FilePath: /allfunc/gin_admin/routers/base.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

import (
	v1 "project/allfunc/gin_admin/api/v1"

	"github.com/gin-gonic/gin"
)

type OrderRouter struct{}

func (oa *OrderRouter) InitOrderRouter(Router *gin.RouterGroup) {

	orderApi := v1.ApiGroupApp.SystemApiGroup.OrderApi
	Router.GET("/order/list", orderApi.OrderList)

}
