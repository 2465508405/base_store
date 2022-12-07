/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-15 17:30:47
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-11-04 16:45:09
 * @FilePath: /allfunc/gin_admin/routers/system/enter.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package system

type RouterGroup struct {
	UserRouter
	HomeRouter
	OrderRouter
	LoginRouter
	GoodsRouter
}
