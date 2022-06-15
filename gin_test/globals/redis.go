/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-05-28 21:04:35
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-05-29 13:16:05
 * @FilePath: /allfunc/gin_test/globals/redis.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package globals

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client

func init() {
	// RedisClient = redis.NewClient(&redis.Options{
	// 	// host:port address.
	// 	Addr:     ":6379",
	// 	Password: "",
	// 	// Database to be selected after connecting to the server.
	// 	DB: 0,
	// })
	// _, _ = RedisClient.Set(context.Background(), "s", "ss", 0).Result()
	//集群模式
	// RedisCluster := redis.NewClusterClient(&redis.ClusterOptions{
	// 	Addrs: []string{"127.0.0.1:6380"},
	// })
	// result, err := RedisCluster.Set(context.Background(), "cluster", "aaa", 0).Result()
	// fmt.Println("redis init", err, result)
	//哨兵模式
	client := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName: "mymaster",
		// A seed list of host:port addresses of sentinel nodes.
		SentinelAddrs: []string{"127.0.0.1:26380", "127.0.0.1:26381", "127.0.0.1:26382"},
	})
	fmt.Println("sss")
	for {
		reply, err := client.Incr(context.Background(), "pvcount").Result()
		fmt.Printf("reply=%v err=%v\n", reply, err)
		time.Sleep(1 * time.Second)
	}
	// fmt.Println("success", RedisRepl)

}
