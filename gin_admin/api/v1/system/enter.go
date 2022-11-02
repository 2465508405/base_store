/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-15 14:52:42
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-11-01 10:30:32
 * @FilePath: /allfunc/gin_admin/api/v1/system/enter.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

import "project/allfunc/gin_admin/service"

type ApiGroup struct {
	UserApi
	HomeApi
	LoginApi
	OrderApi
}

var (
	// apiService              = service.ServiceGroupApp.ApiService
	// jwtService              = service.ServiceGroupApp.JwtService
	// menuService             = service.ServiceGroupApp.SystemServiceGroup.MenuService
	userService = service.ServiceGroupApp.UserService
	// initDBService           = service.ServiceGroupApp.InitDBService
	// casbinService           = service.ServiceGroupApp.CasbinService
	// autoCodeService         = service.ServiceGroupApp.AutoCodeService
	// baseMenuService         = service.ServiceGroupApp.BaseMenuService
	// authorityService        = service.ServiceGroupApp.AuthorityService
	// dictionaryService       = service.ServiceGroupApp.DictionaryService
	// systemConfigService     = service.ServiceGroupApp.SystemConfigService
	// operationRecordService  = service.ServiceGroupApp.OperationRecordService
	// autoCodeHistoryService  = service.ServiceGroupApp.AutoCodeHistoryService
	// dictionaryDetailService = service.ServiceGroupApp.DictionaryDetailService
	// authorityBtnService     = service.ServiceGroupApp.AuthorityBtnService
)
