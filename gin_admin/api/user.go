/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-17 12:53:10
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-07-18 19:00:14
 * @FilePath: /allfunc/leju_test/api/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer33.com!",
	})
}
