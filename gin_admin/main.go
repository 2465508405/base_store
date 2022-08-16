/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-17 11:53:42
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-16 17:23:49
 * @FilePath: /allfunc/leju_test/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"project/allfunc/gin_admin/global"
	"project/allfunc/gin_admin/initialize"
)

func main() {

	global.GVA_DB = initialize.InitDB()

	r := initialize.InitRouter()

	initialize.MigrateTables(global.GVA_DB)
	if err := r.Run(":8000"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
