/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-16 22:57:37
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-11-04 15:14:39
 * @FilePath: /allfunc/gin_admin/service/system/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

import (
	"errors"
	"fmt"
	"project/allfunc/gin_admin/common"
	"project/allfunc/gin_admin/global"
	"project/allfunc/gin_admin/models/system"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// type IUserService interface {
// }

type UserService struct {
	// UserRepository repositories.IUserRepository
}

// func NewService(repository repositories.IUserRepository) {
// return &UserService{repository}
// fmt.Println(repository)
// }

func (us *UserService) UserList() []system.User {
	db := common.NewMysqlConn()
	// db := global.GVA_DB
	var users []system.User
	db.Select("id", "name", "email", "status", "intro", "age").Find(&users)
	return users
}

func (us *UserService) UserRegister(c *gin.Context) error {
	db := common.NewMysqlConn()

	username := c.PostForm("name")
	pwd := c.PostForm("password")

	// hashPass, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	hashPass, err := bcrypt.GenerateFromPassword([]byte(pwd), 10)
	if err != nil {
		return err
	}
	user := system.User{Name: username, Password: string(hashPass)}

	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (us *UserService) UserLogin(c *gin.Context) (system.User, error) {
	db := common.NewMysqlConn()
	name := c.PostForm("name")
	passwod := c.PostForm("password")
	var user system.User
	result := db.Where("name = ?", name).First(&user)
	if result.Error != nil {
		return system.User{}, result.Error
	}
	hashedPass := user.Password
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(passwod)); err != nil {
		return system.User{}, errors.New("密码比对错误！")
	}

	return user, nil
}

func (us *UserService) UserCreate(c *gin.Context) bool {
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

func (u *UserService) UserDel(c *gin.Context) bool {
	id := c.PostForm("id")
	// db := global.GVA_DB
	db := common.NewMysqlConn()
	result := db.Delete(&system.User{}, id)

	return result.RowsAffected > 0
}
