/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-16 22:57:37
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-16 23:00:12
 * @FilePath: /allfunc/gin_admin/service/system/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

import (
	"project/allfunc/gin_admin/global"
	"project/allfunc/gin_admin/models/system"
)

type UserService struct{}

func (us *UserService) UserList() []system.User {
	db := global.GVA_DB
	var users []system.User
	db.Select("id", "name", "email", "status", "intro", "age").Find(&users)
	return users
}
