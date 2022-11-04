/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-11-04 10:23:13
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-11-04 15:28:31
 * @FilePath: /allfunc/gin_admin/lib/redis/redis.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package redisclient

import (
	"fmt"
	"project/allfunc/gin_admin/global"
	"strconv"

	"github.com/go-redis/redis/v8"
)

type RedisServer struct {
	Client *redis.Client
}

func NewRedisServer() (s RedisServer) {

	fmt.Println("addr: port: ", global.GVA_CONFIG.RedisConfig.Host, global.GVA_CONFIG.RedisConfig.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     global.GVA_CONFIG.RedisConfig.Host + ":" + strconv.Itoa(global.GVA_CONFIG.RedisConfig.Port),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	s = RedisServer{Client: client}
	return
}
