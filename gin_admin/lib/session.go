/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-12 17:48:04
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-11-03 15:41:01
 * @FilePath: /allfunc/gin_admin/lib/session.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package lib

import (
	"fmt"
	"net/http"
	"project/allfunc/gin_admin/models/system"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// var store = sessions.NewCookieStore([]byte("user_login_key"))

var store = sessions.NewFilesystemStore("./tmp/sess", securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(32))

func SessionSet(c *gin.Context, user system.User) {

	session, _ := store.Get(c.Request, "session-name")
	// if err != nil {
	// 	panic(err)
	// }
	session.Values["id"] = user.ID
	session.Values["name"] = user.Name
	// 保存更改
	session.Save(c.Request, c.Writer)
	fmt.Println("home:afafaffafafa")
}

func GetSession(r *http.Request, w http.ResponseWriter) *sessions.Session {
	session, err := store.Get(r, "session-name")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil
	}
	foo := session.Values["id"]
	fmt.Println("user_id:", foo)
	return session
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
