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
		fmt.Println("位置：", offset)
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
