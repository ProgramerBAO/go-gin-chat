package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
	"go-gin-chat-server/middleware"
	"net/http"
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// ConnWs SendMsg 群聊广播 返回ws连接
func ConnWs(ctx *gin.Context) *websocket.Conn {
	// 连接服务器
	conn, err := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		fmt.Println("连接服务器异常", err)
		return nil
	}
	return conn
}

// SendMsg 群聊广播
func SendMsg(ctx *gin.Context) {
	conn := ConnWs(ctx)
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Println("关闭失败", err)
		}
	}(conn)
	// 先订阅消息 这里堵住了
	pubSub := middleware.Subscribe(ctx, middleware.PublishKey)
	// 回传数据
	go func(*redis.PubSub) {
		// 往回发
		receiveMessages := pubSub.Channel()
		for message := range receiveMessages {
			err := conn.WriteMessage(1, []byte("收到了"+conn.RemoteAddr().String()+"发送的:"+message.Payload))
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}(pubSub)
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println(string(message))
		err = middleware.Publish(ctx, middleware.PublishKey, string(message))
		if err != nil {
			return
		}
	}
	defer middleware.ClosePubsub(pubSub)
}
