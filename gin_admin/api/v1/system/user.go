/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-17 12:53:10
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-11-01 18:28:05
 * @FilePath: /allfunc/leju_test/api/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

import (
	"fmt"
	"net/http"
	"project/allfunc/gin_admin/global"
	"project/allfunc/gin_admin/models/system"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Title string
	Name  string
	Users []system.User
}
type UserApi struct{}

func (ua *UserApi) UserList(c *gin.Context) {

	users := userService.UserList()
	Info := UserInfo{Title: "后台", Name: "sfafaf", Users: users}

	c.HTML(http.StatusOK, "user/list.html", Info)
}

func (ua *UserApi) UserAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "user/add.html", gin.H{"title": "后台管理系统", "address": "www.5lmh.com"})
}

func (ua *UserApi) UserCreate(c *gin.Context) {
	// db := global.GVA_DB
	// username := c.PostForm("name")
	// password := c.PostForm("password")

	// pass := global.Md5Crypt(password)
	// var userInfo system.User
	// _ = c.ShouldBindJSON(&userInfo)
	// fmt.Println("userInfo:", userInfo)
	if result := userService.UserCreate(c); !result {
		fmt.Println("创建失败")
	}
	fmt.Println("3333")
	// db.Create(&system.User{Name: username, Password: pass})
	c.Redirect(http.StatusSeeOther, "/user/list") //303
	// c.HTML(http.StatusOK, "user/list.html", gin.H{"title": "后台管理系统", "address": "www.5lmh.com"})
}

func (ua *UserApi) UserEdit(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	db := global.GVA_DB
	var user system.User
	db.First(&user, id)

	c.HTML(http.StatusOK, "user/edit.html", gin.H{"title": "后台管理系统", "user": user})
}

func (ua *UserApi) UserUpdate(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	password = global.Md5Crypt(password)
	status, _ := strconv.Atoi(c.PostForm("status"))
	db := global.GVA_DB
	var user system.User
	db.First(&user, id)
	result := db.Model(&user).Updates(system.User{Name: name, Email: email, Status: status, Password: password})
	if result.RowsAffected > 0 {
		c.Redirect(http.StatusTemporaryRedirect, "/user/list")
	}
}

func (ua *UserApi) UserDel(c *gin.Context) {
	var msg string = "删除成功"
	if ok := userService.UserDel(c); !ok {
		msg = "删除失败"
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}

func (ua *UserApi) FileUpload(c *gin.Context) {

}
