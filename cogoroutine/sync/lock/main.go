/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-09-07 14:14:41
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-07 14:27:02
 * @FilePath: /allfunc/cogoroutine/sync/lock/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x       int64
	wg      sync.WaitGroup
	lock    sync.Mutex
	rwLock  sync.RWMutex
	rwLock1 sync.RWMutex
)

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}

	for i := 0; i < 1; i++ {
		wg.Add(1)
		go read()
	}

	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func write() {
	// lock.Lock()   // 加互斥锁
	rwLock.Lock() // 加写锁
	x = x + 1
	fmt.Println("x:", x)
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwLock.Unlock()                   // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwLock.RLock()                       // 加读锁
	time.Sleep(time.Millisecond * 10000) // 假设读操作耗时1毫秒
	rwLock.RUnlock()                     // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg.Done()
}
