/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-09 17:33:29
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-07-09 18:07:31
 * @FilePath: /allfunc/rabbitmq/sendrecv/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func main() {

	dial, err := amqp.Dial("amqp://guest:guest@localhost:5672")

	if err != nil {
		panic(err)
	}

	channel, err := dial.Channel()
	if err != nil {
		panic(err)
	}
	q, err := channel.QueueDeclare(
		"go_q1", //队列名字
		true,    //durable 是否持久化
		false,   //自动删除
		false,   //exlusive 专有的
		false,   // nowait 不等待
		nil,     //
	)
	if err != nil {
		panic(err)
	}
	defer channel.Close()

	go consumer("c1", channel, q)
	go consumer("c2", channel, q)
	i := 1
	for {
		time.Sleep(time.Millisecond * 100)

		fmt.Println("写入", i)

		err = channel.Publish(
			"",     //exchange
			q.Name, //队列名
			false,  //mandatory 强制性的
			false,  //immediate 立即, 消息是否立即被消费
			amqp.Publishing{
				Body: []byte(fmt.Sprintf("这里是写入队列中的信息 %d", i)),
			},
		)

		i++

		if err != nil {
			fmt.Println(err.Error())
		}
	}

}

func consumer(consumer string, channel *amqp.Channel, q amqp.Queue) {
	go func() {
		consume, err := channel.Consume(
			q.Name,   //队列
			consumer, //消费名
			true,     //autoack
			false,    //exlusive 专有的
			false,    //nolocal
			false,    //nowait
			nil,      //args
		)

		if err != nil {
			panic(err)
		}
		for msg := range consume {
			fmt.Printf("%s:%s\n", consumer, string(msg.Body))
		}

	}()
}
