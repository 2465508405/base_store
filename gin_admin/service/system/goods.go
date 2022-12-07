/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-11-04 16:56:33
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-11-04 16:59:36
 * @FilePath: /allfunc/gin_admin/service/system/goods.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

import (
	"fmt"
	"project/allfunc/gin_admin/common"
	"project/allfunc/gin_admin/global"
	"project/allfunc/gin_admin/models/system"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GoodsService struct {
}

func (g *GoodsService) GoodsCreate(c *gin.Context) bool {
	db := common.NewMysqlConn()
	// db := global.GVA_DB

	username := c.PostForm("name")
	password := c.PostForm("password")
	email := c.PostForm("email")
	status, _ := strconv.Atoi(c.PostForm("status"))
	pass := global.Md5Crypt(password)

	user := system.User{Name: username, Password: pass, Email: email, Status: status}

	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	}
	fmt.Println("user:id:", user.ID)
	return true
}

func (g *GoodsService) GoodsUpate(c *gin.Context) bool {
	db := common.NewMysqlConn()
	// db := global.GVA_DB

	username := c.PostForm("name")
	// password := c.PostForm("password")
	// email := c.PostForm("email")
	// status, _ := strconv.Atoi(c.PostForm("status"))
	// pass := global.Md5Crypt(password)

	user := system.Goods{Name: username}

	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return false
	}
	fmt.Println("user:id:", user.ID)
	return true
}

func (g *GoodsService) GoodsDel(c *gin.Context) bool {
	id := c.PostForm("id")
	// db := global.GVA_DB
	db := common.NewMysqlConn()
	result := db.Delete(&system.User{}, id)

	return result.RowsAffected > 0
}
