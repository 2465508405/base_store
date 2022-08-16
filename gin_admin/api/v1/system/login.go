/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-12 10:03:29
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-16 17:40:15
 * @FilePath: /allfunc/gin_admin/api/v1/system/login.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

import (
	"fmt"
	"net/http"
	"project/allfunc/gin_admin/global"
	"project/allfunc/gin_admin/models/system"

	"github.com/gin-gonic/gin"
)

type LoginApi struct{}

func (la *LoginApi) Login(c *gin.Context) {

	c.HTML(http.StatusOK, "login/login.html", gin.H{"title": "后台管理系统", "address": "www.5lmh.com"})

}

func (la *LoginApi) AuthLogin(c *gin.Context) {
	name := c.PostForm("name")
	passwod := c.PostForm("password")
	var user system.User
	result := global.GVA_DB.Where("name = ? and password = ? ", name, global.Md5Crypt(passwod)).First(&user)
	fmt.Println(result)
	cookie, err := c.Cookie("sessionid")
	if err != nil {
		cookie = "NotSet"
		// 给客户端设置cookie
		//  maxAge int, 单位为秒
		// path,cookie所在目录
		// domain string,域名
		//   secure 是否智能通过https访问
		// httpOnly bool  是否允许别人通过js获取自己的cookie
		c.SetCookie("sessionid", "5", 60, "/",
			"localhost", false, true)
	}
	fmt.Printf("cookie的值是： %s\n", cookie)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"msg": "登陆成功", "code": 0, "address": "www.5lmh.com"})
	}
}
