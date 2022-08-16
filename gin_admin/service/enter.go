/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-15 18:26:50
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-15 18:28:19
 * @FilePath: /allfunc/gin_admin/service/enter.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import "project/allfunc/gin_admin/service/system"

type ServiceGroup struct {
	SystemServiceGroup system.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
