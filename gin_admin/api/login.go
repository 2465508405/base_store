package api

import (
	"fmt"
	"net/http"
	"project/allfunc/gin_admin/global"
	"project/allfunc/gin_admin/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	c.HTML(http.StatusOK, "login/login.html", gin.H{"title": "后台管理系统", "address": "www.5lmh.com"})

}

func AuthLogin(c *gin.Context) {
	name := c.PostForm("name")
	passwod := c.PostForm("password")
	var user models.User
	result := global.DB.Where("name = ? and password = ? ", name, global.Md5Crypt(passwod)).First(&user)
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
