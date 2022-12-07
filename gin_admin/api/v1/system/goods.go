/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-11-04 15:51:24
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-11-04 16:52:06
 * @FilePath: /allfunc/gin_admin/api/v1/system/goods.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GoodsApi struct {
}

func (g *GoodsApi) GoodsLists(c *gin.Context) {

	c.HTML(http.StatusOK, "goods/index.html", gin.H{"title": "后台管理系统", "address": "www.5lmh.com"})
}

func (g *GoodsApi) GoodsCreate(c *gin.Context) {

}

func (g *GoodsApi) GoodsUpdate(c *gin.Context) {

}

func (g *GoodsApi) GoodsDel(c *gin.Context) {

}
