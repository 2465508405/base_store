/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-05-28 14:34:11
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-05-28 14:34:52
 * @FilePath: /allfunc/redis/bloom/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"

	"./data"

	"github.com/go-redis/redis/v8"
)

func main() {

	con := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", "127.0.0.1", 6380),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	// con, err := redis.Dial("tcp", ":6379") //连接redis
	// print(err, "connect")
	defer con.Close()

	bloom := data.NewBloom(con)   //创建过滤器
	bloom.Add("aa")               //往过滤器写入数据
	b := bloom.Exist("newClient") //判断是否存在这个值
	fmt.Println(b)
	// // Connect to localhost with no password
	// var client = redisbloom.NewClient("localhost:6381", "nohelp", nil)

	// // BF.ADD mytest item
	// _, err := client.Add("mytest", "myItem")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }

	// exists, err := client.Exists("mytest", "myItem")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// }
	// fmt.Println("myItem exists in mytest: ", exists)
}
