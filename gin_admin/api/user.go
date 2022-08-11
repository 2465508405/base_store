/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-17 12:53:10
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-11 14:38:11
 * @FilePath: /allfunc/leju_test/api/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package api

import (
	"net/http"
	"project/allfunc/gin_admin/global"
	"project/allfunc/gin_admin/models"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello www.topgoer33.com!",
	})
}

func UserList(c *gin.Context) {

	c.HTML(http.StatusOK, "user/list.html", gin.H{"title": "后台管理系统", "address": "www.5lmh.com"})
}
func UserAdd(c *gin.Context) {
	c.HTML(http.StatusOK, "user/add.html", gin.H{"title": "后台管理系统", "address": "www.5lmh.com"})
}

func UserCreate(c *gin.Context) {
	db := global.DB
	username := c.PostForm("name")
	password := c.PostForm("password")

	db.Create(&models.User{Username: username, Password: password})
	c.HTML(http.StatusOK, "user/list.html", gin.H{"title": "后台管理系统", "address": "www.5lmh.com"})
}

func UserEdit(c *gin.Context) {
	c.HTML(http.StatusOK, "user/edit.html", gin.H{"title": "后台管理系统", "address": "www.5lmh.com"})
}

func UserUpdate(c *gin.Context) {

}

func UserDel(c *gin.Context) {

}

func FileUpload(c *gin.Context) {

}
