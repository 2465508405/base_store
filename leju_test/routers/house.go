/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-17 13:12:15
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-07-18 18:58:25
 * @FilePath: /allfunc/leju_test/routers/house.go
 * @Description:
 */
package routers

import (
	"project/allfunc/leju_test/api"

	"github.com/gin-gonic/gin"
)

func InitHouse(r *gin.Engine) {

	r.GET("/house/list", api.GetHouseList)
}
