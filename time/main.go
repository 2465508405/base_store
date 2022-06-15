/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-05-22 18:37:00
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-05-22 19:24:04
 * @FilePath: /allfunc/time/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type JsonTime time.Time

func (jt JsonTime) MarshalJSON() ([]byte, error) {
	var stmp = fmt.Sprintf("\"%s\"", time.Time(jt).Format("2006-01-02"))
	return []byte(stmp), nil
}

type UserResponse struct {
	Birthday JsonTime `json:"birthday"`
}

func main() {

	timestamp := time.Now().Unix() - 1000 //获取当前时间time时间戳

	ur := UserResponse{
		Birthday: JsonTime(time.Now()),
	}

	js, _ := json.Marshal(ur)
	fmt.Println(string(js))       //生成日期时间json
	year := time.Now().Year()     //获取当前年份
	month := time.Now().Month()   //获取当前月份
	day := time.Now().Day()       //获取当前天
	hour := time.Now().Hour()     //获取当前小时
	minute := time.Now().Minute() //获取当前分钟
	second := time.Now().Second() //获取当前秒

	time.Sleep(10 * time.Second) //休眠10秒s

	timeStr := time.Now().Format("2006-01-02 15:04:05") //获取当前时间年月日
	//时间戳转日期格式
	timeDate := time.Unix(timestamp, 0).Format("2006-01-02 15:04:05") //时间戳转化成日期格式

	fmt.Println(timeDate)
	fmt.Println(timestamp)
	fmt.Println(timeStr)
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)
	fmt.Println(hour)
	fmt.Println(minute)
	fmt.Println(second)
}
