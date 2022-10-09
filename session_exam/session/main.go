/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-19 14:29:50
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-20 16:00:00
 * @FilePath: /allfunc/session_exam/session/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// Note: Don't store your key in your source code. Pass it via an
// environmental variable, or flag (or both), and don't accidentally commit it
// alongside your code. Ensure your key is sufficiently random - i.e. use Go's
// crypto/rand or securecookie.GenerateRandomKey(32) and persist the result.
var store = sessions.NewCookieStore([]byte("123abc"))

func SessionSet(w http.ResponseWriter, r *http.Request) {
	// Get a session. We're ignoring the error resulted from decoding an
	// existing session: Get() always returns a session, even if empty.
	session, _ := store.Get(r, "session-name")
	// Set some session values.
	session.Values["foo"] = "bar"
	session.Values[42] = 43
	fmt.Println("777")
	// Save it before we write to the response/return from the handler.
	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func SessionGet(w http.ResponseWriter, r *http.Request) {
	sess, err := store.Get(r, "session-name")

	fmt.Println(sess.Values["foo"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {

	http.HandleFunc("/session/set", SessionSet)
	http.HandleFunc("/session/get", SessionGet)
	http.ListenAndServe(":9090", nil)
}
