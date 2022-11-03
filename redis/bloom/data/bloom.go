/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-05-28 14:34:11
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-11-03 17:19:47
 * @FilePath: /allfunc/redis/bloom/data/bloom.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package data

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Bloom struct {
	Conn      *redis.Client
	Key       string
	HashFuncs []F //保存hash函数
}

func NewBloom(con *redis.Client) *Bloom {
	return &Bloom{Conn: con, Key: "bloom", HashFuncs: NewFunc()}
}

func (b *Bloom) Add(str string) error {
	var err error
	for _, f := range b.HashFuncs {
		offset := f(str)
		fmt.Println(offset)
		err := b.Conn.SetBit(context.Background(), b.Key, int64(offset), 1).Err()
		if err != nil {
			return err
		}
	}
	return err
}
func (b *Bloom) Exist(str string) bool {
	var a int64 = 1
	for _, f := range b.HashFuncs {
		offset := f(str)
		bitValue, _ := b.Conn.GetBit(context.Background(), b.Key, int64(offset)).Result()
		if bitValue != a {
			return false
		}
	}
	return true
}
