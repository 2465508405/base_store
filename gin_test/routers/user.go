/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-05-28 16:34:32
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-05-28 16:36:43
 * @FilePath: /allfunc/gin_test/router/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package routers

import "github.com/gin-gonic/gin"

func GetUserList(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "get user list",
	})
}
