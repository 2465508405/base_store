/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-10-28 14:35:25
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-10-28 17:02:25
 * @FilePath: /allfunc/gin_admin/repositories/user_repository.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE

 */
package repositories

import (
	"project/allfunc/gin_admin/initialize"

	"gorm.io/gorm"
)

type IUserRepository interface {
	Conn() error
}

type UserManagerRepository struct {
	mysqlConn *gorm.DB
}

func NewUserRepository(table string, db *gorm.DB) IUserRepository {
	return &UserManagerRepository{db}
}

func (u *UserManagerRepository) Conn() (err error) {
	if u.mysqlConn == nil {
		mysql := initialize.NewMysqlConn()

		u.mysqlConn = mysql
	}

	return
}

func (u *UserManagerRepository) GetUserList() (err error) {

	return
}
