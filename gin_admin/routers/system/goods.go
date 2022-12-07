/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-11-04 16:40:59
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-11-04 16:52:00
 * @FilePath: /allfunc/gin_admin/routers/system/goods.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

import (
	v1 "project/allfunc/gin_admin/api/v1"

	"github.com/gin-gonic/gin"
)

type GoodsRouter struct {
}

func (g *GoodsRouter) InitGoodsRouter(r *gin.RouterGroup) {

	goodsApi := v1.ApiGroupApp.SystemApiGroup.GoodsApi

	r.GET("/login", goodsApi.GoodsLists)
	r.GET("/register", goodsApi.GoodsCreate)
	r.POST("/auth/login", goodsApi.GoodsUpdate)
	r.POST("/auth/register", goodsApi.GoodsDel)
}
