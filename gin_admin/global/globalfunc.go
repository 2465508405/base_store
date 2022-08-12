/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-12 10:28:26
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-12 10:28:38
 * @FilePath: /allfunc/gin_admin/global/globalfunc.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package global

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func Md5Crypt(password string) string {
	m := md5.New()
	io.WriteString(m, password)
	pass := hex.EncodeToString(m.Sum(nil))
	return pass
}
