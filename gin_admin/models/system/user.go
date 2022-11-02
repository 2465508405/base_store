/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-10 17:08:36
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-11-01 10:44:02
 * @FilePath: /allfunc/gin_admin/models/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name" gorm:"type:varchar(32);not null;default:'';index:idx_name"`
	Mobile   string `json:"mobile" gorm:"type:varchar(11);"`
	Email    string `json:"email" gorm:"type:varchar(32);not null;default:'';"`
	Avatar   string `json:"avatar" gorm:"type:varchar(128);not null;default:'';"`
	Role     string `json:"role" gorm:"type:varchar(32);"`
	Age      int    `json:"age" gorm:"type:tinyint(4);"`
	Birth    string `json:"birth" gorm:"type:varchar(16);"`
	Sex      string `json:"sex" gorm:"type:enum('FEMALE', 'MALE', 'UNKNOWN');default:UNKNOWN;"`
	Password string `json:"password" gorm:"type:varchar(128);"`
	Token    string `json:"token" gorm:"type:varchar(128);not null;default:'';comment:token验证值;"`
	Status   int    `json:"status" gorm:"type:tinyint(2);not null;default:1;comment:用户状态：1.启用 0.禁用;"`
	Intro    string `json:"intro" gorm:"type:varchar(255);not null;default:'';comment:介绍"`
}
