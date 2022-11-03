/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-12 10:03:29
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-11-03 15:53:37
 * @FilePath: /allfunc/gin_admin/api/v1/system/login.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

import (
	"fmt"
	"net/http"
	"project/allfunc/gin_admin/lib/encrypt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LoginApi struct{}

func (la *LoginApi) Login(c *gin.Context) {

	c.HTML(http.StatusOK, "login/login.html", gin.H{"title": "后台管理系统", "address": "www.5lmh.com"})

}

func (la *LoginApi) Register(c *gin.Context) {
	c.HTML(http.StatusOK, "login/register.html", gin.H{"title": "后台管理系统", "address": "www.5lmh.com"})

}

func (la *LoginApi) AuthRegister(c *gin.Context) {

	if err := userService.UserRegister(c); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
			"data": map[string]interface{}{
				"msg": "err",
			},
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"error": 200,
		"data": map[string]interface{}{
			"msg": "注册成功",
		},
	})

}

func (la *LoginApi) AuthLogin(c *gin.Context) {

	user, result := userService.UserLogin(c)
	if result != nil {
		fmt.Println("err: login->", result)
		c.JSON(http.StatusOK, gin.H{"msg": "登陆失败", "code": 0, "address": "www.5lmh.com"})
		return
	}
	// cookie, err := c.Cookie("sessionid")
	// if err != nil {
	// 给客户端设置cookie
	//  maxAge int, 单位为秒
	// path,cookie所在目录
	// domain string,域名
	//   secure 是否智能通过https访问
	// httpOnly bool  是否允许别人通过js获取自己的cookie
	useridbyte := strconv.Itoa(int(user.ID))
	c.SetCookie("sessionid", useridbyte, 0, "/",
		"localhost", false, true)
	uidByte := []byte(useridbyte)
	uidString, err := encrypt.EnPwdCode(uidByte)
	if err != nil {
		fmt.Println(err)
	}
	//写入用户浏览器
	c.SetCookie("sign", uidString, 0, "/", "localhost", false, true)

	// global.GVA_USER = user
	// lib.SessionSet(c, user) //设置session

	// lib.GetSession(c.Request, c.Writer) // 获取session
	// }
	// fmt.Printf("cookie的值是： %s\n", cookie)
	// if result.RowsAffected > 0 {
	c.JSON(http.StatusOK, gin.H{"msg": "登陆成功", "code": 0, "address": "www.5lmh.com"})
	// }
}
