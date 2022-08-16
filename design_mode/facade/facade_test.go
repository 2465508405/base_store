/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-16 15:35:05
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-08-16 15:35:08
 * @FilePath: /allfunc/design_mode/facade/facade_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package facade

import "testing"

var expect = "A module running\nB module running"

// TestFacadeAPI ...
func TestFacadeAPI(t *testing.T) {
	api := NewAPI()
	ret := api.Test()
	if ret != expect {
		t.Fatalf("expect %s, return %s", expect, ret)
	}
}
