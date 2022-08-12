/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-10 10:56:50
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-12 15:21:17
 * @FilePath: /allfunc/gin_admin/middlewares/Login.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package middlewares

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/yudeguang/slice"
)

func LoginMiddleWare(routers []interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		path := c.Request.RequestURI

		fmt.Printf("urlss :%s\n", path)
		fmt.Printf("time:%s\n", t)
		if slice.Contains(routers, path) {
			c.Next()
			return
		}
		fmt.Println("home:afafaffafafa")
		if cookie, err := c.Cookie("sessionid"); err == nil {
			fmt.Printf("get login cookie info : %s\n", cookie)
			if cookie == "5" {
				c.Next()
				return
			}
		}
		fmt.Println("中间件开始执行了")
		// 返回错误
		c.Redirect(http.StatusTemporaryRedirect, "/login")
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}
