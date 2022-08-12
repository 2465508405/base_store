/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-17 12:53:10
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-12 11:45:41
 * @FilePath: /allfunc/leju_test/api/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package api

import (
	"net/http"
	"project/allfunc/gin_admin/global"
	"project/allfunc/gin_admin/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Title string
	Name  string
	Users []models.User
}

func UserList(c *gin.Context) {
	db := global.DB
	var users []models.User
	db.Select("id", "name", "email", "status", "intro", "age").Find(&users)
	Info := UserInfo{Title: "后台", Name: "sfafaf", Users: users}

	c.HTML(http.StatusOK, "user/list.html", Info)
}

func UserAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "user/add.html", gin.H{"title": "后台管理系统", "address": "www.5lmh.com"})
}

func UserCreate(c *gin.Context) {
	db := global.DB
	username := c.PostForm("name")
	password := c.PostForm("password")

	pass := global.Md5Crypt(password)

	db.Create(&models.User{Name: username, Password: pass})
	c.Redirect(http.StatusMovedPermanently, "/user/list")
	// c.HTML(http.StatusOK, "user/list.html", gin.H{"title": "后台管理系统", "address": "www.5lmh.com"})
}

func UserEdit(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	db := global.DB
	var user models.User
	db.First(&user, id)

	c.HTML(http.StatusOK, "user/edit.html", gin.H{"title": "后台管理系统", "user": user})
}

func UserUpdate(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	email := c.PostForm("email")
	password := c.PostForm("password")
	password = global.Md5Crypt(password)
	status, _ := strconv.Atoi(c.PostForm("status"))
	db := global.DB
	var user models.User
	db.First(&user, id)
	result := db.Model(&user).Updates(models.User{Name: name, Email: email, Status: status, Password: password})
	if result.RowsAffected > 0 {
		c.Redirect(http.StatusMovedPermanently, "/user/list")
	}
}

func UserDel(c *gin.Context) {
	id := c.PostForm("id")
	db := global.DB
	result := db.Delete(&models.User{}, id)
	if result.RowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg": "删除成功",
		})
	}
}

func FileUpload(c *gin.Context) {

}
