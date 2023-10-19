package middleware

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go-gin-chat-server/utils"
)

// PublishKey 这里指的是频道 就好像Welcome to my channel
const PublishKey = "websocket"

// Publish 发布消息到redis channel 是频道名字
func Publish(ctx context.Context, channel, msg string) error {
	res := utils.RDB.Publish(ctx, channel, msg)
	// fmt.Println(channel)
	if res.Err() != nil {
		fmt.Println("出错啦:", res.Err())
		return res.Err()
	}
	fmt.Println("发布成功!消息是:", res.Val())
	fmt.Println("发布成功!消息是:", res)
	return nil
}

// Subscribe 群聊  消息取
func Subscribe(ctx context.Context, channel string) *redis.PubSub {
	sub := utils.RDB.Subscribe(ctx, channel)
	fmt.Println(channel)
	//defer func(sub *redis.PubSub) {
	//	err := sub.Close()
	//	if err != nil {
	//		fmt.Println("关闭失败!!!!")
	//	}
	//	fmt.Println("关闭成功!!!")
	//}(sub)
	fmt.Println("show:", sub)
	fmt.Println("show:", ctx)

	// 连接时长
	//for message := range ch {
	//	fmt.Println(message.Channel, message.Payload)
	//}
	// 得到了管道
	//var msg *redis.Message
	//var err error
	//for {
	//	msg, err = sub.ReceiveMessage(ctx)
	//	if err != nil {
	//		fmt.Println("出错啦:", err)
	//		break
	//	}
	//	fmt.Println("收到消息,正在回传:", msg.Payload)
	//}
	//// 转换字符串
	//fmt.Println("收到消息,正在回传:", msg.Payload)
	// return msg.Payload, nil

	return sub
}

func ClosePubsub(sub *redis.PubSub) {
	err := sub.Close()
	if err != nil {
		fmt.Println("关闭失败!!!!")
	}
	fmt.Println("关闭成功!!!")
}
