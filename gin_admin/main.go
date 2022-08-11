/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-17 11:53:42
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-10 22:31:18
 * @FilePath: /allfunc/leju_test/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"project/allfunc/gin_admin/initialize"
	"project/allfunc/gin_admin/routers"
)

func main() {

	// info := map[string]interface{}{"name": "ykk", "age": 13}

	// for key, val := range info {
	// 	fmt.Println("key, val,", key, val)
	// }

	initialize.InitDB()
	r := routers.InitRouter()

	if err := r.Run(":8000"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
