/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-08-15 14:16:49
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-11-02 17:49:06
 * @FilePath: /test_info/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"errors"
	"fmt"
	"sync"
	// "github.com/robfig/cron/v3"
)

type Parent struct {
	Stu
}
type Stu struct {
	Id  int
	Age string
}

func (s *Stu) add() {
	fmt.Println("aasss")
}

var h int

var mp []int

type sa struct {
}
type (
	cc Stu
)

var wg sync.WaitGroup

func Greeting(prefix string, who ...string) {
	if who == nil {
		fmt.Printf("Nobody to say hi.")
		return
	}

	for _, people := range who {
		fmt.Printf("%s %s\n", prefix, people)
	}
}
func Validation() []error {
	var errs []error

	errs = append(errs, errors.New("error 1"))
	errs = append(errs, errors.New("error 2"))
	errs = append(errs, errors.New("error 3"))

	return errs
}

type (
// 计划任务
// Cron = cron.New()
)

func main() {
	c1 := new(cc)
	c1.Id = 7
	c1.Age = "haha"
	fmt.Println(c1)
}

func IsPanic() bool {
	if err := recover(); err != nil {
		fmt.Println("Recover success...")
		return true
	}

	return false
}

var (
	name int
	age  string
)

func UpdateTable() {
	// defer中决定提交还是回滚
	defer func() {
		if IsPanic() {

		}
	}()
	panic("ahha")

	// Database update operation...
}
