/*
 * @Author: ykk ykk@qq.com
 * @Date: 2022-07-09 17:33:29
 * @LastEditors: ykk ykk@qq.com
 * @LastEditTime: 2022-09-06 10:45:52
 * @FilePath: /allfunc/rabbitmq/sendrecv/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

// const exchange = "exchange_fanout"

const exchange = "exchange_topic"

// const exchange = "exchange_direct"

func main() {

	// direct_mode()
	// fanout_mode()
	topic_mode()
}

func direct_mode() {
	dial, err := amqp.Dial("amqp://guest:guest@localhost:5672")

	if err != nil {
		panic(err)
	}

	channel, err := dial.Channel()
	if err != nil {
		panic(err)
	}

	if err := channel.ExchangeDeclare(
		exchange,
		"direct", //类型 fanout, topic, direct
		true,     //durable 持久化
		true,     //autodelete 自动删除
		false,    // internal 内置交换机
		false,    //nowait
		nil,      //args
	); err != nil {
		log.Fatalf("cannot declare fanout exchange: %v", err)
	}

	q1, err := channel.QueueDeclare(
		"dir_que1", //队列名字
		true,       //durable 是否持久化
		true,       //自动删除
		false,      //exlusive 专有的， //独占队列只能由声明它们和的连接访问将在连接关闭时删除。其他连接上的通道在试图声明、绑定、使用、清除或删除同名队列会报错。
		true,       // nowait 不等待
		nil,        //
	)
	// q2, err := channel.QueueDeclare(
	// 	"",    //队列名字
	// 	true,  //durable 是否持久化
	// 	false, //自动删除
	// 	false, //exlusive 专有的
	// 	false, // nowait 不等待
	// 	nil,   //
	// )

	if err := channel.QueueBind(
		q1.Name,       //队列名字
		"routingkey1", // 路由key
		exchange,      // 交换机名子
		false,         //notwait
		nil,           //args 参数
	); err != nil {
		log.Printf("cannot consume without a binding to exchange: %q, %v", exchange, err)
		return
	}

	// // if err := channel.QueueBind(
	// // 	q2.Name,       //队列名字
	// // 	"routingkey2", // 路由key
	// // 	exchange,      // 交换机名子
	// // 	false,         //notwait
	// // 	nil,           //args 参数
	// // ); err != nil {
	// // 	log.Printf("cannot consume without a binding to exchange: %q, %v", exchange, err)
	// // 	return
	// // }

	// if err != nil {
	// 	panic(err)
	// }
	// defer channel.Close()

	// go consumer_direct("c1", channel, q1)
	// // go consumer("c2", channel, q2)
	// i := 1
	for i := 0; i < 10; i++ {
		time.Sleep(time.Millisecond * 100)

		fmt.Println("写入", i)
		//通过fanout, topic, direct机制 ，交换机将数据推送到队列中， 不需要填写队列名

		//1、mandatory：当mandatory设置为true时，如果exchange根据自身类型和消息routeKey无法找到一个符合条件的queue，那么会调用basic.return方法将消息返还给生产者；当mandatory设为false时，出现上述情形broker会直接将消息扔掉。
		//2、immediate：当immediate设置为true时，如果exchange在将消息route到queue时发现对应的queue上没有消费者，那么这条消息不会放入队列中。当与消息routeKey关联的所有queue(一个或多个)都没有消费者时，该消息会通过basic.return方法返还给生产者。
		//3、DeliveryMode: 投递模式分Transient（临时的） 或者 Persistent（永久的）。设置为Persistent时，消息会持久化到磁盘。

		err = channel.Publish(
			exchange,        //exchange
			"routingkey122", //路由参数
			false,           //mandatory 强制性的
			false,           //immediate 立即, 消息是否立即被消费
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				Body:         []byte(fmt.Sprintf("这里是写入队列中的信息 %d", i)),
			},
		)

		// 	i++

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func fanout_mode() {

	dial, err := amqp.Dial("amqp://guest:guest@localhost:5672")

	if err != nil {
		panic(err)
	}

	channel, err := dial.Channel()
	if err != nil {
		panic(err)
	}

	if err := channel.ExchangeDeclare(
		exchange,
		"fanout", //类型 fanout, topic, direct
		false,    //durable 持久化
		true,     //autodelete 自动删除
		false,    // internal 计时
		false,    //nowait
		nil,      //args
	); err != nil {
		log.Fatalf("cannot declare fanout exchange: %v", err)
	}

	q1, err := channel.QueueDeclare(
		"",    //队列名字 未设置， 默认会自动创建key
		true,  //durable 是否持久化
		false, //自动删除
		false, //exlusive 专有的
		false, // nowait 不等待
		nil,   //
	)
	q2, err := channel.QueueDeclare(
		"",    //队列名字
		true,  //durable 是否持久化
		false, //自动删除
		false, //exlusive 专有的
		false, // nowait 不等待
		nil,   //
	)

	if err := channel.QueueBind(
		q1.Name,  //队列名字
		"",       // 路由key
		exchange, // 交换机名子
		false,    //notwait
		nil,      //args 参数
	); err != nil {
		log.Printf("cannot consume without a binding to exchange: %q, %v", exchange, err)
		return
	}

	if err := channel.QueueBind(
		q2.Name,  //队列名字
		"",       // 路由key
		exchange, // 交换机名子
		false,    //notwait
		nil,      //args 参数
	); err != nil {
		log.Printf("cannot consume without a binding to exchange: %q, %v", exchange, err)
		return
	}

	if err != nil {
		panic(err)
	}
	defer channel.Close()

	go consumer_fanout("c1", channel, q1)
	go consumer_fanout("c2", channel, q2)
	i := 1
	for {
		time.Sleep(time.Millisecond * 100)

		fmt.Println("写入", i)
		//通过fanout, topic, direct机制 ，交换机将数据推送到队列中， 不需要填写队列名
		err = channel.Publish(
			exchange, //exchange
			"",       //队列名
			false,    //mandatory 强制性的
			false,    //immediate 立即, 消息是否立即被消费
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
func topic_mode() {

	dial, err := amqp.Dial("amqp://guest:guest@localhost:5672")

	if err != nil {
		panic(err)
	}

	channel, err := dial.Channel()
	if err != nil {
		panic(err)
	}

	if err := channel.ExchangeDeclare(
		exchange,
		"topic", //类型 fanout, topic, direct
		false,   //durable 持久化
		true,    //autodelete 会否自动删除  ,设置为true ,不存在队列或交换机绑定时，会自动删除
		false,   // internal 是否是内置交换机
		false,   //nowait 是否等待服务器确认
		nil,     //args 其它配置
	); err != nil {
		log.Fatalf("cannot declare fanout exchange: %v", err)
	}
	exchange1 := "exch1_topic"
	if err := channel.ExchangeDeclare(
		exchange1,
		"topic", //类型 fanout, topic, direct
		false,   //durable 持久化
		true,    //autodelete 会否自动删除  ,设置为true ,不存在队列或交换机绑定时，会自动删除
		false,   // internal 是否是内置交换机
		false,   //nowait 是否等待服务器确认
		nil,     //args 其它配置
	); err != nil {
		log.Fatalf("cannot declare fanout exchange: %v", err)
	}

	q1, err := channel.QueueDeclare(
		"",    //队列名字
		true,  //durable 是否持久化
		true,  //自动删除 设置为true 队列数据执行完之后，会将队列删除，false 则保留
		false, //exlusive 专有的
		false, // nowait 不等待
		nil,   //
	)
	// q2, err := channel.QueueDeclare(
	// 	"",    //队列名字
	// 	true,  //durable 是否持久化
	// 	true,  //自动删除
	// 	false, //exlusive 专有的
	// 	false, // nowait 不等待
	// 	nil,   //
	// )

	if err := channel.QueueBind(
		q1.Name,  //队列名字
		"*.com",  // 路由key
		exchange, // 交换机名子
		false,    //notwait
		nil,      //args 参数
	); err != nil {
		log.Printf("cannot consume without a binding to exchange: %q, %v", exchange, err)
		return
	}

	// if err := channel.QueueBind(
	// 	q2.Name,  //队列名字
	// 	"*.com",  // 路由key
	// 	exchange, // 交换机名子
	// 	false,    //notwait
	// 	nil,      //args 参数
	// ); err != nil {
	// 	log.Printf("cannot consume without a binding to exchange: %q, %v", exchange, err)
	// 	return
	// }

	// err = channel.ExchangeBind(exchange, "routingkey.com", exchange1, false, nil)

	if err != nil {
		panic(err)
	}
	defer channel.Close()

	go consumer_topic("c1", channel, q1)
	// go consumer_topic("c2", channel, q2)
	i := 1
	for {
		time.Sleep(time.Millisecond * 200)

		fmt.Println("写入", i)
		//通过fanout, topic, direct机制 ，交换机将数据推送到队列中， 不需要填写队列名
		err = channel.Publish(
			// exchange1,        //exchange
			exchange,
			"routingkey.com", //RouterKey
			false,            //mandatory 是否为无法路由的消息进行返回处理
			false,            //immediate 否对路由到无消费者队列的消息进行返回处理 RabbitMQ 3.0 废弃
			amqp.Publishing{
				// DeliveryMode: amqp.Persistent, //Msg set as persistent
				Body: []byte(fmt.Sprintf("这里是写入队列中的信息 %d", i)),
			},
		)

		i++

		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func consumer_direct(consumer string, channel *amqp.Channel, q amqp.Queue) {
	go func() {
		consume, err := channel.Consume( //推模式
			q.Name,   //队列
			consumer, //消费名 消费者名称，自定义，可以为空。
			false,    //autoack
			false,    //exlusive 专有的 设置是否排他，排他表示当前队列只能给一个消费者使用。
			false,    //nolocal 如果为true，表示生产者和消费者不能是同一个connect。
			false,    //nowait  ：是否非阻塞，true表示是。阻塞：表示创建交换器的请求发送后，阻塞等待RMQ Server返回信息。非阻塞：不会阻塞等待RMQ Server的返回信息，而RMQ Server也不会返回信息
			nil,      //args
		)

		// msg , ok ,err := channel.Get(q.Name, true)//拉模式

		if err != nil {
			panic(err)
		}
		for msg := range consume {
			if err := msg.Ack(true); err != nil {
				fmt.Printf("消息确认失败：%s\n", err.Error())
			} //被动确认

			msg.Reject(true) //如果requeue为true，则RMQ会把这条消息重新加入队列，如果requeue为false，则RMQ会丢弃本条消息。
			fmt.Printf("%s:%s\n", consumer, string(msg.Body))
		}

	}()
}

func consumer_fanout(consumer string, channel *amqp.Channel, q amqp.Queue) {
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

func consumer_topic(consumer string, channel *amqp.Channel, q amqp.Queue) {

	consume, err := channel.Consume(
		q.Name,   //队列
		consumer, //消费名
		false,    //autoack 是否确认消费 设置false需要 ack确认
		false,    //exlusive 专有的 排他
		false,    //nolocal
		false,    //nowait
		nil,      //args
	)
	// consume, ok, err := channel.Get(q.Name, false) //拉模式
	// if err != nil {
	// 	panic(err)
	// }
	// if ok {
	// 	fmt.Printf("%s:%s", consumer, consume)
	// }

	if err != nil {
		panic(err)
	}
	for msg := range consume {
		// if err := msg.Ack(true); err != nil {
		// 	fmt.Println(err.Error())
		// }
		msg.Ack(true) //确认机制
		// channel.Ack(msg.DeliveryTag, false)
		fmt.Printf("%s:%s\n", consumer, string(msg.Body))

	}
}
