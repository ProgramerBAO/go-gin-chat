package servicemodels

import (
	"github.com/fatih/set"
	"github.com/gorilla/websocket"
)

// Node 的意思是存放一个连接,这样就方便建立通讯,
// 思考: 建立通讯池会不会搞高效些
type Node struct {
	UserId    int64
	Conn      *websocket.Conn
	DataQueue chan []byte
	// github.com/fatih/set 这个包用于对集合进行操作 取交集差集
	// 可以用map[string]Type 代替
	GroupSets set.Interface
}
