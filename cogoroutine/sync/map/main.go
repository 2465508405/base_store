/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-07 14:39:53
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-07 15:16:39
 * @FilePath: /allfunc/cogoroutine/sync/map/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"strconv"
	"sync"
)

var m = make(map[string]int)

var lock sync.RWMutex

func get(key string) int {
	return m[key]
}

func set(key string, value int) {
	lock.Lock()
	m[key] = value
	lock.Unlock()
}

//Go语言中内置的map不是并发安全的, 并发执行过程中，会相互影响，处理过程中逻辑的判断
//if h.flags&hashWriting != 0 { // h是指向map,
// throw("concurrent map writes")
//}
func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(n int) {
			key := strconv.Itoa(n)
			fmt.Println(key)
			set(key, n)
			fmt.Printf("k=:%v,v:=%v\n", key, get(key))
			wg.Done()
		}(i)
	}
	wg.Wait()
}
