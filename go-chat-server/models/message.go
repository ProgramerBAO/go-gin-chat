package models

import (
	"gorm.io/gorm"
)

// Message 这里放消息的结构 传递消息 这里只放对于数据库的操作
// 特别是非文字  图片 文件 之类的 可以存放在一个地方然后用链接形式 小而美 节省空间
type Message struct {
	gorm.Model
	FromId   int64  //发送者
	TargetId int64  //接受者
	Type     int    // 消息模式 群聊 私聊 广播
	Media    int    // 消息类型 文字 图片 广播 音频
	Content  string // 消息内容
	Pic      string // 图片消息
	Url      string // 链接消息
	Desc     string // 描述相关
	Amount   int    // 其他的数据统计
}

func (table *Message) TableName() string {
	return "message"
}

// Node 的意思是存放一个连接,这样就方便建立通讯,
// 思考: 建立通讯池会不会搞高效些
//type Node struct {
//	Conn      *websocket.Conn
//	DataQueue chan []byte
//	// github.com/fatih/set 这个包用于对集合进行操作 取交集差集
//	// 可以用map[string]Type 代替
//	GroupSets set.Interface
//}

// 映射关系 这里就是存储连接的地方
// var clientMap map[int64]*Node = make(map[int64]*Node)
//var clientMap sync.Map

// 读写锁 因为读写map所以加锁
//var rwLocker sync.RWMutex

// Chat 需要: 发送者ID,接受者ID,消息类型 发送的内容 发送类型
//func Chat(writer http.ResponseWriter, request *http.Request) {
//	// 获取参数并检验token 还没做
//	// token := query.Get("token")
//	// fmt.Println("连接成功")
//	// 这里获取到相关参数 但是并没有在建立连接的时候给出来?
//	query := request.URL.Query()
//	userId := query.Get("userId")
//	// userId应该是唯一的建立连接
//	userIdInt, err2 := strconv.ParseInt(userId, 10, 64)
//	if err2 != nil {
//		fmt.Println(err2)
//		// 因为没有传入所以是0
//		userIdInt = 0
//	}
//	//fmt.Println("连接成功")
//	//msgType := query.Get("msgType")
//	//targetId := query.Get("targetId")
//	//context := query.Get("context")
//	isValid := true // 检查token
//	conn, err := (&websocket.Upgrader{
//		// token 校验
//		// 这里校验的是请求头,在后续的开发中 请求头会带有验证信息,通过下面的方法实现鉴权
//		CheckOrigin: func(r *http.Request) bool {
//			return isValid
//		},
//	}).Upgrade(writer, request, nil)
//	// 以上我们完成了ws的链接 以下我们得到了地址和端口
//	// remote 是客户端的ip
//	fmt.Println(conn.RemoteAddr().Network())
//	fmt.Println(conn.RemoteAddr().String())
//	// 本地是8080 也就是服务端
//	fmt.Println(conn.LocalAddr().String())
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	// 获取连接
//	node := &Node{
//		Conn:      conn,
//		DataQueue: make(chan []byte, 50),
//		// 创建一个线程安全的散列表
//		GroupSets: set.New(set.ThreadSafe),
//	}
//	// 用户关系 先不搞
//
//	//userid跟node绑定 加锁 这样的话就相当于入池子了
//	// 比如说服务器是一个聊天室,那么用户进来了
//	//rwLocker.Lock()
//	//clientMap[userIdInt] = node
//	//rwLocker.Unlock()
//	clientMap.Store(userIdInt, node)
//	// 完成发送的逻辑 就是开始实时准备发送消息
//	go sendProc(node)
//	// 完成接收的逻辑
//	// 实时准备接收消息
//	go receProc(node)
//	sendMsg(userIdInt, userIdInt, []byte("欢迎"))
//}
//
//// sendProc 发送消息
//func sendProc(node *Node) {
//	for {
//		select {
//		case data := <-node.DataQueue:
//			// 这里只是文本? 写出去 但是到那里呢?
//			err := node.Conn.WriteMessage(websocket.TextMessage, data)
//			if err != nil {
//				fmt.Println(err)
//				return
//			}
//		}
//	}
//}
//
//// receProc 接收消息
//func receProc(node *Node) {
//	for {
//		_, message, err := node.Conn.ReadMessage()
//		if err != nil {
//			return
//		}
//		// 广播消息 ? 不是私聊吗?
//		broadMsg(message)
//		fmt.Println("ws<<<<<<<", string(message))
//	}
//}
//
//// 用管道存数据
//var udpSendChan chan []byte = make(chan []byte, 1024)
//
//func broadMsg(message []byte) {
//	udpSendChan <- message
//}
//
//func init() {
//	// udp 数据发送协程
//	go udpSendProc()
//	// udp 数据接收协程
//	go udpReceProc()
//}
//
//func udpSendProc() {
//	// 在这里发送个给指定的人?
//	udpConn, err := net.DialUDP("udp", nil, &net.UDPAddr{
//		IP:   net.IPv4(127, 0, 0, 1),
//		Port: 3000,
//	})
//
//	defer closeUDPConn(udpConn)
//
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	for {
//		select {
//		case data := <-udpSendChan:
//			_, err := udpConn.Write(data)
//			if err != nil {
//				return
//			}
//		}
//	}
//
//}
//
//func udpReceProc() {
//	udpConn, err := net.ListenUDP("udp", &net.UDPAddr{
//		IP:   net.IPv4zero,
//		Port: 3000,
//	})
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	defer closeUDPConn(udpConn)
//
//	for {
//		var buf [512]byte
//		read, err := udpConn.Read(buf[0:])
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//		dispatch(buf[0:read])
//	}
//}
//
//// 后端调度逻辑处理
//func dispatch(bytes []byte) {
//	msg := Message{}
//	err := json.Unmarshal(bytes, &msg)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	// 根据类型确定信息类型
//	switch msg.Type {
//	// 这里建议使用常量 不然是魔法值
//	case 1:
//		// 私聊
//		sendMsg(msg.FromId, msg.TargetId, bytes)
//		//case 2:
//		//	// 群聊
//		//	sendGroupMsg()
//		//case 3:
//		//	// 广播
//		//	sendAllMsg()
//		//case 4:
//
//	}
//}
//
//func sendMsg(fromId, targetId int64, msg []byte) {
//	//rwLocker.RLock()
//	//node, ok := clientMap[targetId]
//	//rwLocker.RUnlock()
//	node, ok := clientMap.Load(targetId)
//	if ok {
//		node.DataQueue <- msg
//	}
//}
//
//// 关闭UDP连接
//func closeUDPConn(udpConn *net.UDPConn) {
//	err := udpConn.Close()
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//}
