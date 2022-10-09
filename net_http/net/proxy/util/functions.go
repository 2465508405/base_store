/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-15 22:48:58
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-15 23:01:05
 * @FilePath: /allfunc/net_http/net/proxy/util/functions.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package util

import "net/http"

func CloneHeader(src http.Header, dst *http.Header) {
	for k, v := range src {
		dst.Set(k, v[0])
	}
}
