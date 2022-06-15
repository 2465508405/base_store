/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-05-24 11:50:13
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-05-28 14:25:55
 * @FilePath: /allfunc/redis/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

//字符串操作  结构类型：int, embstr, raw
func StrOp() {
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}
	_ = rdb.SetNX(ctx, "key", "value", 0).Err()
	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	//将给定 key 的值设为 value ，并返回 key 的旧值(old value)。
	oldval, _ := rdb.GetSet(ctx, "key", "newvalue").Result()
	fmt.Println("oldvalue : ", oldval)
	//对 key 所储存的字符串值，获取指定偏移量上的位(bit)。
	bitval, _ := rdb.GetBit(ctx, "key", 2).Result()
	fmt.Println("bitvalu :", bitval)
	//获取所有(一个或多个)给定 key 的值。
	mval, _ := rdb.MGet(ctx, "key", "incr1").Result()
	fmt.Println("mget: ", mval)
	//对 key 所储存的字符串值，设置或清除指定偏移量上的位(bit)。
	bitset, _ := rdb.SetBit(ctx, "key", 1, 1).Result()
	fmt.Println("set bit :", bitset)
	//将值 value 关联到 key ，并将 key 的过期时间设为 seconds (以秒为单位)。
	err = rdb.SetEX(ctx, "exKey", 10, 100).Err()
	if err != nil {
		panic(err)
	}
	exVal, _ := rdb.Get(ctx, "exKey").Result()
	fmt.Println("expire val: ", exVal)
	//只有在 key 不存在时设置 key 的值。
	_ = rdb.SetNX(ctx, "nxKey", "nx", 0).Err()
	nxVal, _ := rdb.Get(ctx, "exKey").Result()
	fmt.Println("nx val: ", nxVal)
	//用 value 参数覆写给定 key 所储存的字符串值，从偏移量 offset 开始。 返回字符串长度
	rangeVal, _ := rdb.SetRange(ctx, "nxKey", 1, "hhh").Result()
	fmt.Println("setrang val:", rangeVal)
	//返回 key 所储存的字符串值的长度。
	length, _ := rdb.StrLen(ctx, "nxKey").Result()
	fmt.Println("nxkey len:", length)
	//
	//返回 key 中字符串值的子字符
	substr, _ := rdb.GetRange(ctx, "key", 1, -1).Result()
	fmt.Println("subvalu: ", substr)
	//同时设置一个或多个 key-value 对，当且仅当所有给定 key 都不存在。
	err = rdb.MSetNX(ctx, "key1", "val1", "key2", "val2").Err()
	if err != nil {
		fmt.Println("setnx val:", err)
	}
	//如果 key 已经存在并且是一个字符串， APPEND 命令将指定的 value 追加到该 key 原来值（value）的末尾。 返回字符串长度
	len, _ := rdb.Append(ctx, "key1", "append").Result()
	fmt.Println("append len:", len)
	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	//key是否存在
	num, err := rdb.Exists(ctx, "key2").Result()
	if err != nil {
		panic(err)
	}
	if num == 0 {
		fmt.Println("key exists")
	}
	//incr 增加  返回递增的值
	sum, err := rdb.Incr(ctx, "incr1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(sum)

	//incrBy
	sum, _ = rdb.IncrBy(ctx, "incr1", 3).Result()
	fmt.Println(sum)
	//decr
	sum, _ = rdb.Decr(ctx, "incr1").Result()
	fmt.Println(sum)
	//decrby
	sum, _ = rdb.DecrBy(ctx, "incr1", 3).Result()
	fmt.Println(sum)

	//keys *
	keys, _ := rdb.Keys(ctx, "*").Result()
	fmt.Println(keys)

	//del
	del, _ := rdb.Del(ctx, "key").Result()
	if del > 0 {
		fmt.Println("success, ", del)
	}

}

//hash函数, ziplist(包含上一条记录的长度，和当前数据的长度), hashtable
func HashOp() {
	//将哈希表 key 中的字段 field 的值设为 value 。
	exist, _ := rdb.HSet(ctx, "hkey1", "name", "ykk", "age", 10, "birth", "2020-01-11").Result()
	if exist > 0 {
		fmt.Println("hset val :", exist)
	}
	//获取存储在哈希表中指定字段的值。返回删除的数量
	info, _ := rdb.HGet(ctx, "hkey1", "name").Result()
	fmt.Println("hget vale:", info)
	//获取所有给定字段的值
	all, _ := rdb.HMGet(ctx, "hkey1", "name", "age").Result()
	fmt.Println("hmget val:", all)
	//删除一个或多个哈希表字段,返回删除的数量
	del, _ := rdb.HDel(ctx, "hkey1", "name").Result()
	if del > 0 {
		fmt.Println("del: ", del)
	}
	//查看哈希表 key 中，指定的字段是否存在。
	exis, _ := rdb.HExists(ctx, "hkey1", "age").Result()
	if !exis {
		fmt.Println("not exit hkey1 name")
	} else {
		fmt.Println("exists hkey1 name")
	}
	//获取在哈希表中指定 key 的所有字段和值
	hall, _ := rdb.HGetAll(ctx, "hkey1").Result()
	fmt.Println("hget all: ", hall)
	//为哈希表 key 中的指定字段的整数值加上增量 increment 。
	incr, _ := rdb.HIncrBy(ctx, "hkey1", "age", 2).Result()
	fmt.Println("hincr val:", incr)
	//获取所有哈希表中的字段
	seg, _ := rdb.HKeys(ctx, "hkey1").Result()
	fmt.Println("all keys :", seg)
	//获取哈希表中字段的数量
	len, _ := rdb.HLen(ctx, "hkey1").Result()
	fmt.Println("key len :", len)
	//获取哈希表中所有值。
	vals, _ := rdb.HVals(ctx, "hkey1").Result()
	fmt.Println("hkey len:", vals)

	//只有在字段 field 不存在时，设置哈希表字段的值。
	es, _ := rdb.HSetNX(ctx, "hkey1", "name", "ykk").Result()
	if es {
		fmt.Println(" setnx val:", es)
	}
}

//列表操作   ziplist, quicklist
func ListOp() {
	//将一个或多个值插入到列表头部, 返回数据条数
	lp, _ := rdb.LPush(ctx, "lkey1", 1, "3", "ssf", "si", "wu", "liu").Result()
	fmt.Println("lpush :", lp)
	//移出并获取列表的第一个元素
	lpop, _ := rdb.LPop(ctx, "lkey1").Result()
	fmt.Println("lpop :", lpop)
	//	获取列表长度
	len, _ := rdb.LLen(ctx, "lkey1").Result()
	fmt.Println("llen :", len)
	//获取列表指定范围内的元素
	lrange, _ := rdb.LRange(ctx, "lkey1", 0, -1).Result()
	fmt.Println("lrange :", lrange)
	//移除列表元素,返回移除元素个数
	lrem, _ := rdb.LRem(ctx, "lkey1", 2, "3").Result()
	fmt.Println("rem count:", lrem)
	lrang, _ := rdb.LRange(ctx, "lkey1", 0, -1).Result()
	fmt.Println("lrange :", lrang)
	//通过索引设置列表元素的值
	lset, _ := rdb.LSet(ctx, "lkey1", 2, "haha").Result()
	fmt.Println("lset :", lset)
	lran, _ := rdb.LRange(ctx, "lkey1", 0, -1).Result()
	fmt.Println("lrange :", lran)
	//对一个列表进行修剪(trim)，就是说，让列表只保留指定区间内的元素，不在指定区间之内的元素都将被删除。
	ltrim, _ := rdb.LTrim(ctx, "lkey1", 0, 1).Result()
	fmt.Println("ltrim :", ltrim)
	lra, _ := rdb.LRange(ctx, "lkey1", 0, -1).Result()
	fmt.Println("lrange :", lra)
	//移除列表的最后一个元素，返回值为移除的元素。
	rpop, _ := rdb.RPop(ctx, "lkey1").Result()
	fmt.Println("rpop :", rpop)
	lr, _ := rdb.LRange(ctx, "lkey1", 0, -1).Result()
	fmt.Println("lrange :", lr)
	//移除列表的最后一个元素，并将该元素添加到另一个列表并返回
	lpoppush, _ := rdb.RPopLPush(ctx, "lkey1", "lkey2").Result()
	fmt.Println("rpopLpush :", lpoppush)
	//在列表中添加一个或多个值
	rp, _ := rdb.RPush(ctx, "lkey1", 1, "3", "ssf", "si", "wu", "liu").Result()
	fmt.Println("lpush :", rp)

	//为已存在的列表添加值,返回列表长度
	pushex, _ := rdb.RPushX(ctx, "lkey1", "12", "ssss").Result()
	fmt.Println("pushex :", pushex)
	//移出并获取列表的第一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
	blpop, _ := rdb.BLPop(ctx, time.Second*10, "lkey1").Result()
	fmt.Println("blpop :", blpop)
	//移出并获取列表的最后一个元素， 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
	brpop, _ := rdb.BRPop(ctx, time.Second*10, "lkey1").Result()
	fmt.Println("brpop :", brpop)
	//从列表中弹出一个值，将弹出的元素插入到另外一个列表中并返回它； 如果列表没有元素会阻塞列表直到等待超时或发现可弹出元素为止。
	brpoplpush, _ := rdb.BRPopLPush(ctx, "lkey1", "lkey2", time.Second*100).Result()
	fmt.Println("brpoplpush :", brpoplpush)
	//通过索引获取列表中的元素
	lindex, _ := rdb.LIndex(ctx, "lkey1", 2).Result()
	fmt.Println("lindex :", lindex)
	//在列表的元素前或者后插入元素,成功返回列表长度 before | after
	linsert, _ := rdb.LInsert(ctx, "lkey1", "before", "wu", "cccgggg").Result()
	fmt.Println("linsert :", linsert)

}

//集合操作  intset, hashtable
func SetOp() {
	//向集合添加一个或多个成员
	sadd, _ := rdb.SAdd(ctx, "setkey1", "a", "b", "c", "d").Result()
	fmt.Println("sadd val:", sadd)
	sadd1, _ := rdb.SAdd(ctx, "setkey2", "a", "b", "c", "e", "f").Result()
	fmt.Println("sadd1 val:", sadd1)
	//返回第一个集合与其他集合之间的差异。
	sdiff, _ := rdb.SDiff(ctx, "setkey1", "setkey2").Result()
	fmt.Println("sdiff val: ", sdiff)
	//返回给定所有集合的差集并存储在 destination 中
	sstore, _ := rdb.SDiffStore(ctx, "setkey3", "setkey1", "setkey2").Result()
	fmt.Println("sstore diffStore :", sstore)
	//返回给定所有集合的交集
	sinter, _ := rdb.SInter(ctx, "setkey1", "setkey2").Result()
	fmt.Println("sinter val:", sinter)
	//返回给定所有集合的交集并存储在 destination 中
	sinterstore, _ := rdb.SInterStore(ctx, "setkey4", "setkey1", "setkey2").Result()
	fmt.Println("sstore diffStore :", sinterstore)
	//判断 member 元素是否是集合 key 的成员
	smem, _ := rdb.SIsMember(ctx, "setkey1", "a").Result()
	fmt.Println("sismemver val:", smem)
	//返回集合中的所有成员
	smems, _ := rdb.SMembers(ctx, "setkey1").Result()
	fmt.Println("smembers val:", smems)
	//将 member 元素从 source 集合移动到 destination 集合
	smove, _ := rdb.SMove(ctx, "setkey1", "setkeyMove", "a").Result()
	fmt.Println("smove:", smove)
	//移除并返回集合中的一个随机元素
	spop, _ := rdb.SPop(ctx, "setkey1").Result()
	fmt.Println("spop val:", spop)
	//返回集合中一个或多个随机数
	srandmem, _ := rdb.SRandMember(ctx, "setkey1").Result()
	fmt.Println("srandmember vale:", srandmem)
	//移除集合中一个或多个成员
	srem, _ := rdb.SRem(ctx, "setkey1", "a").Result()
	fmt.Println("srem val:", srem)
	//返回所有给定集合的并集
	sunion, _ := rdb.SUnion(ctx, "setkey1", "setkey2").Result()
	fmt.Println("sunion val:", sunion)
	//所有给定集合的并集存储在 destination 集合中
	sunionstore, _ := rdb.SUnionStore(ctx, "setkeyUnion", "setkey1", "setkey2").Result()
	fmt.Println("sunionstore val:", sunionstore)
	//迭代集合中的元素
	// scan, _ := rdb.SScan(ctx, "setkey1", 1, "b", 10).Result()
	// fmt.Println("scan val:", scan)
}

//有序集合 ziplist, skiplist
func zsetOp() {
	//向有序集合添加一个或多个成员，或者更新已存在成员的分数
	zadd, _ := rdb.ZAdd(ctx, "zkey1", &redis.Z{Score: 1, Member: "a"}, &redis.Z{Score: 2, Member: "b"}, &redis.Z{Score: 3, Member: "c"}, &redis.Z{Score: 4, Member: "d"}).Result()
	fmt.Println("zadd val:", zadd)
	zadd1, _ := rdb.ZAdd(ctx, "zkey2", &redis.Z{Score: 1, Member: "b"}, &redis.Z{Score: 2, Member: "c"}, &redis.Z{Score: 3, Member: "d"}, &redis.Z{Score: 4, Member: "e"}).Result()
	fmt.Println("zadd1 val:", zadd1)
	//获取有序集合的成员数
	zcard, _ := rdb.ZCard(ctx, "zkey1").Result()
	fmt.Println("zcard val:", zcard)
	//计算在有序集合中指定区间分数的成员数
	zcount, _ := rdb.ZCount(ctx, "zkey1", "2", "4").Result()
	fmt.Println("zcount val:", zcount)
	//有序集合中对指定成员的分数加上增量 increment
	zincrby, _ := rdb.ZIncrBy(ctx, "zkey1", 2, "a").Result()
	fmt.Println("zincrby val:", zincrby)
	//计算给定的一个或多个有序集的交集并将结果集存储在新的有序集合 destination 中
	zinterStore, _ := rdb.ZInterStore(ctx, "zinterStore", &redis.ZStore{
		Keys:    []string{"zkey1", "zkey2"},
		Weights: []float64{2, 3},
	}).Result()
	fmt.Println("zinterStore val:", zinterStore)
	//在有序集合中计算指定字典区间内成员数量["-", "+"] ["a", "b"]
	zlexcount, _ := rdb.ZLexCount(ctx, "zkey1", "-", "+").Result()
	fmt.Println("zlexcount val:", zlexcount)
	//通过索引区间返回有序集合指定区间内的成员
	zrange, _ := rdb.ZRange(ctx, "zkey1", 0, -1).Result()
	fmt.Println("zrange val:", zrange)
	//通过字典区间返回有序集合的成员
	zrangebylex, _ := rdb.ZRangeByLex(ctx, "zkey1", &redis.ZRangeBy{
		Min:    "a",
		Max:    "d",
		Offset: 1,
		Count:  10,
	}).Result()
	fmt.Println("zrangebylex val:", zrangebylex)
	//通过分数返回有序集合指定区间内的成员
	zrangebyScore, _ := rdb.ZRangeByScore(ctx, "zkey1", &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  2,
	}).Result()
	fmt.Println("zrangebyScore val:", zrangebyScore)
	//返回有序集合中指定成员的索引
	zrank, _ := rdb.ZRank(ctx, "zkey1", "a").Result()
	fmt.Println("zrank val:", zrank)
	//移除有序集合中的一个或多个成员
	zrem, _ := rdb.ZRem(ctx, "zkey1", "b").Result()
	fmt.Println("zrem val:", zrem)
	//移除有序集合中给定的字典区间的所有成员
	zremrangebylex, _ := rdb.ZRemRangeByLex(ctx, "zkey1", "b", "d").Result()
	fmt.Println("zremrangebylex val:", zremrangebylex)
	//移除有序集合中给定的排名区间的所有成员
	zremrangebyRank, _ := rdb.ZRemRangeByRank(ctx, "zkey1", 2, 4).Result()
	fmt.Println("zremrangebyScore val:", zremrangebyRank)
	//移除有序集合中给定的分数区间的所有成员
	zremrangebyScore, _ := rdb.ZRemRangeByScore(ctx, "zkey1", "30", "40").Result()
	fmt.Println("zremrangebyscore val:", zremrangebyScore)
	//返回有序集中指定区间内的成员，通过索引，分数从高到低
	zrevrange, _ := rdb.ZRevRange(ctx, "zkey1", 1, 5).Result()
	fmt.Println("zrevrange val:", zrevrange)
	//返回有序集中指定分数区间内的成员，分数从高到低排序
	zrevrangebyscore, _ := rdb.ZRevRangeByScore(ctx, "zkey1", &redis.ZRangeBy{
		Min:    "b",
		Max:    "d",
		Offset: 0,
		Count:  2,
	}).Result()
	fmt.Println("zrevrangebyscore", zrevrangebyscore)
	//返回有序集合中指定成员的排名，有序集成员按分数值递减(从大到小)排序
	zrevrank, _ := rdb.ZRevRank(ctx, "zkey1", "d").Result()
	fmt.Println("zrevrang val:", zrevrank)
	//返回有序集中，成员的分数值
	zscore, _ := rdb.ZScore(ctx, "zkey1", "a").Result()
	fmt.Println("zscore val:", zscore)
	//计算给定的一个或多个有序集的并集，并存储在新的 key 中
	zunionstore, _ := rdb.ZUnionStore(ctx, "zunionStore", &redis.ZStore{
		Keys:    []string{"zkey1", "zkey2"},
		Weights: []float64{2, 3},
		// Can be SUM, MIN or MAX.
		Aggregate: "sum",
	}).Result()
	fmt.Println("zunionstore val:", zunionstore)
	//迭代有序集合中的元素（包括元素成员和元素分值）
	// zscan, _ := rdb.ZScan(ctx, "zkey1", 2, "d", 3).Result()
}

//地理坐标计算
func geoOp() {
	//添加地理位置的坐标
	geoadd, _ := rdb.GeoAdd(ctx, "geokey1", &redis.GeoLocation{
		Name:      "Catania",
		Longitude: 15.087269,
		Latitude:  37.502669,
	}, &redis.GeoLocation{
		Name:      "Palermo",
		Longitude: 13.361389,
		Latitude:  38.115556,
	}).Result()
	fmt.Println("geoadd :", geoadd)
	//获取地理位置的坐标。
	geopos, _ := rdb.GeoPos(ctx, "geokey1", "Catania", "Palermo").Result()
	fmt.Println("geopos val:", geopos)
	//计算两个位置之间的距离。
	geodist, _ := rdb.GeoDist(ctx, "geokey1", "Catania", "Palermo", "km").Result()
	fmt.Println("geodist val:", geodist)
	//根据用户给定的经纬度坐标来获取指定范围内的地理位置集合。
	georadius, _ := rdb.GeoRadius(ctx, "geokey1", 15, 17, &redis.GeoRadiusQuery{
		Radius: 40,
		Sort:   "ASC",
	}).Result()
	fmt.Println("georadius val:", georadius)
	//根据储存在位置集合里面的某个地点获取指定范围内的地理位置集合。
	georadiusbymember, _ := rdb.GeoRadiusByMember(ctx, "geokey1", "Catania", &redis.GeoRadiusQuery{
		Radius: 4,
		Sort:   "ASC",
	}).Result()
	fmt.Println("georadiusbymember val:", georadiusbymember)
	//返回一个或多个位置对象的 geohash 值。
	geohash, _ := rdb.GeoHash(ctx, "geokey1", "Catania", "Palermo").Result()
	fmt.Println("geohash: ", geohash)
}

//HyperLogLog 是用来做基数统计的算法
func HyperLogOp() {
	//添加指定元素到 HyperLogLog 中。
	pfadd, _ := rdb.PFAdd(ctx, "pfkey1", 1, 3, 4).Result()
	pfadd1, _ := rdb.PFAdd(ctx, "pfkey2", 1, 3, 4, 6, 7, 9).Result()
	fmt.Println("pfadd val", pfadd)
	fmt.Println("pfadd1 val", pfadd1)
	//返回给定 HyperLogLog 的基数估算值。
	pfcount, _ := rdb.PFCount(ctx, "pfkey1").Result()
	fmt.Println("pfcount val", pfcount)
	//将多个 HyperLogLog 合并为一个 HyperLogLog
	pfmerge, _ := rdb.PFMerge(ctx, "pfmerge", "pfkey1", "pfkey2").Result()
	fmt.Println("pfmerge val:", pfmerge)
}

//bitmap操作
func bitmapOp() {
	//设置bit值
	bitset, _ := rdb.SetBit(ctx, "bitkey1", 0, 1).Result()
	_, _ = rdb.SetBit(ctx, "bitkey1", 1, 1).Result()
	fmt.Println("bitset val:", bitset)
	//获取bit某个位置的值
	bitget, _ := rdb.GetBit(ctx, "bitkey1", 0).Result()
	fmt.Println("bitget val:", bitget)
	//统计bit格式
	bitcount, _ := rdb.BitCount(ctx, "bitkey1", &redis.BitCount{
		Start: 0,
		End:   -1,
	}).Result()
	fmt.Println("bitcount val:", bitcount)
}

//事务操作  watch（监听), multi(事务)， exec(执行)
func multiOp() {
	// 监视watch_count的值，并在值不变的前提下将其值+1
	// 事务函数
	key := "watch_count"
	maxRetries := 3000
	increment := func(key string) error {
		// 事务函数
		txf := func(tx *redis.Tx) error {
			// 获得key的当前值或零值
			n, err := tx.Get(ctx, key).Int()
			if err != nil && err != redis.Nil {
				return err
			}
			// 实际的操作代码（乐观锁定中的本地操作）
			n++

			// 操作仅在 Watch 的 Key 没发生变化的情况下提交,
			_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
				pipe.Set(ctx, key, n, 0)
				// time.Sleep(time.Second * 15)
				return nil
			})
			return err
		}
		// err := rdb.Watch(ctx, txf, key)
		// 最多重试 maxRetries 次
		for i := 0; i < maxRetries; i++ {
			err := rdb.Watch(ctx, txf, key)
			if err == nil {
				// 成功
				return nil
			}
			if err == redis.TxFailedErr {
				// 乐观锁丢失 重试
				continue
			}
			// 返回其他的错误
			return err
		}
		// return err
		return errors.New("increment reached maximum number of retries")
	}
	// err := increment(key)
	// if err != nil {
	// 	fmt.Println("err :", err)
	// }

	// 模拟 routineCount 个并发同时去修改 counter3 的值
	routineCount := 10
	var wg sync.WaitGroup
	wg.Add(routineCount)
	for i := 0; i < routineCount; i++ {
		go func() {
			defer wg.Done()
			if err := increment(key); err != nil {
				fmt.Println("increment error:", err)
			}
		}()
	}
	wg.Wait()

	n, err := rdb.Get(context.TODO(), key).Int()
	fmt.Println("ended with", n, err)
}

//lua script 原子性操作
func LuaOp() {

	// luaScript := `
	// if redis.call("GET", KEYS[1]) ~= false then
	// 	return {KEYS[1],"==>",redis.call("get", KEYS[1])}
	// end
	// return false`
	// info, err := rdb.Eval(ctx, luaScript, []string{"lua1", "lua2"}, "lua", "l2").Result()
	// if err != nil {
	// 	fmt.Println("err val:", err)
	// }
	// fmt.Println("info val:", info)

	EchoKey := redis.NewScript(`
		if redis.call("GET", KEYS[1]) ~= false then
			return {KEYS[1],"==>",redis.call("get", KEYS[1])}
		end
		return false
	`)

	val1, err := EchoKey.Run(ctx, rdb, []string{"a"}).Result()
	log.Println(val1, err)

	// Lua脚本定义2. 传递key与step使得，key值等于`键值+step`
	IncrByXX := redis.NewScript(`
		if redis.call("GET", KEYS[1]) ~= false then
			return redis.call("INCRBY", KEYS[1], ARGV[1])
		end
		return false
	`)
	// 首次调用
	val2, err := IncrByXX.Run(ctx, rdb, []string{"xx_counter"}, 2).Result()
	log.Println("首次调用 IncrByXX.Run ->", val2, err)

	//加锁，解锁
	str := "abcd"
	LockOp(str)
	UnlockOp(str)
}

//加锁
func LockOp(str string) {
	EchoKey := redis.NewScript(`
		if redis.call('get', KEYS[1]) then 
			return 0;
		else 
			redis.call('setnx', KEYS[1], ARGV[1]);
			redis.call('expire', KEYS[1], ARGV[2]);
			return 1;
		end
	`)
	val1, err := EchoKey.Run(ctx, rdb, []string{"anx"}, str, 30).Result()
	fmt.Println(val1, err)

}

//解锁
func UnlockOp(val string) {
	EchoKey := redis.NewScript(`
		local v = redis.call('get',KEYS[1]);
		if v then 
			-- 如果和传入的值不同，返回0表示失败
			if v~=ARGV[1] then 
				return 0;
			end;
			-- 删除key
			redis.call('del',KEYS[1]);
		end;
		return 1;
	`)
	val1, err := EchoKey.Run(ctx, rdb, []string{"anx"}, val).Result()
	fmt.Println("777")
	fmt.Println(val1, err)
}

func main() {

	// StrOp() //字符串操作
	// HashOp() //hash操作
	// ListOp() //list操作
	// SetOp() //集合操作
	// zsetOp() //有序集合操作
	// geoOp() //地理坐标计算
	// HyperLogOp() //HyperLogLog 是用来做基数统计的算法
	// bitmapOp() //bitmap操作
	// multiOp() //事务操作
	LuaOp() //lua脚本原子性操作
}
