/*
 * @Author: ykk ykk@qq.com
 * @Date: 2023-02-06 16:45:43
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2023-02-06 16:46:00
 * @FilePath: /allfunc/test_info/test/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AEf
 */
package main

import (
	"fmt"
	"runtime"
)

func main() {

	sys := runtime.GOOS

	fmt.Println(sys)
}
