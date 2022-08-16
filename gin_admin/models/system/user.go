/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-10 17:08:36
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-16 16:40:30
 * @FilePath: /allfunc/gin_admin/models/user.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(32);not null;default:'';index:idx_name"`
	Mobile   string `gorm:"type:varchar(11);"`
	Email    string `gorm:"type:varchar(32);not null;default:'';"`
	Avatar   string `gorm:"type:varchar(128);not null;default:'';"`
	Role     string `gorm:"type:varchar(32);"`
	Age      int    `gorm:"type:tinyint(4);"`
	Birth    string `gorm:"type:varchar(16);"`
	Sex      string `gorm:"type:enum('FEMALE', 'MALE', 'UNKNOWN');default:UNKNOWN;"`
	Password string `gorm:"type:varchar(128);"`
	Token    string `gorm:"type:varchar(128);not null;default:'';comment:token验证值;"`
	Status   int    `gorm:"type:tinyint(2);not null;default:1;comment:用户状态：1.启用 0.禁用;"`
	Intro    string `gorm:"type:varchar(255);not null;default:'';comment:介绍"`
}
