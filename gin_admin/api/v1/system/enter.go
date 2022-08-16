/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-15 14:52:42
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-15 15:14:28
 * @FilePath: /allfunc/gin_admin/api/v1/system/enter.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

type ApiGroup struct {
	UserApi
	HomeApi
	LoginApi
	OrderApi
}
