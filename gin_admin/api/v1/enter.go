/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-15 14:50:03
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-17 22:54:33
 * @FilePath: /allfunc/gin_admin/api/enter.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"project/allfunc/gin_admin/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
