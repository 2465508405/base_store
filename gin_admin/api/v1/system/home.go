/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-10 14:10:56
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-11-03 15:54:17
 * @FilePath: /allfunc/gin_admin/api/home.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

import (
	"fmt"
	"net/http"
	"project/allfunc/gin_admin/global"

	"github.com/gin-gonic/gin"
)

type HomeApi struct {
}

func (ua *HomeApi) Home(c *gin.Context) {
	fmt.Println(global.GVA_CONFIG)
	// session := lib.GetSession(c.Request, c.Writer) // 获取session
	// fmt.Println("global user ====>", global.GVA_USER)
	name := c.DefaultQuery("name", "枯藤")
	// fmt.Println("global.GVA_DB:", global.GVA_DB)
	fmt.Println(name)
	c.HTML(http.StatusOK, "home/index.html", gin.H{"title": "后台管理系统", "address": "www.5lmh.com"})
}
