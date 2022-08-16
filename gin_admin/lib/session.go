/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-12 17:48:04
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-12 17:48:28
 * @FilePath: /allfunc/gin_admin/lib/session.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package lib

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("user_login_key"))

func SessionSet(c *gin.Context) {

	session, err := store.Get(c.Request, "session-name")
	if err != nil {
		panic(err)
	}
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	// 保存更改
	session.Save(c.Request, c.Writer)
	fmt.Println("home:afafaffafafa")
}

func GetSession(r *http.Request, w http.ResponseWriter) {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	foo := session.Values["foo"]
	fmt.Println(foo)
}

func SessionExpire(c *gin.Context) {
	session, err := store.Get(c.Request, "session-name")
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
}
