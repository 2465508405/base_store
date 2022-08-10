/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-10 10:56:50
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-10 10:59:47
 * @FilePath: /allfunc/gin_admin/routers/house.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-17 13:12:15
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-10 10:59:03
 * @FilePath: /allfunc/leju_test/routers/house.go
 * @Description:
 */
package routers

import (
	"project/allfunc/gin_admin/api"

	"github.com/gin-gonic/gin"
)

func InitHouse(r *gin.Engine) {

	r.GET("/house/list", api.GetHouseList)
}
