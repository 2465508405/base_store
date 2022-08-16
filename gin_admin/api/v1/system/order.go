/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-13 16:25:01
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-16 16:53:50
 * @FilePath: /allfunc/gin_admin/api/order.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

import (
	"net/http"
	"project/allfunc/gin_admin/global"
	"project/allfunc/gin_admin/models/system"

	"github.com/gin-gonic/gin"
)

type OrderApi struct{}

func (oa *OrderApi) OrderList(c *gin.Context) {
	db := global.GVA_DB
	var orders []system.Orders
	db.Select("id", "name", "price", "order_num", "status", "good_id", "desc").Find(&orders)
	Info := map[string]interface{}{"Title": "后台", "Orders": orders}

	c.HTML(http.StatusOK, "order/list.html", Info)
}
