package servicemodels

import "sync"

// 这里存放需要的结构体

// ClientMap 代表键类型为int、值类型为string的并发安全字典。 存放客户端连接信息
type ClientMap struct {
	m sync.Map
}

func (clientMap *ClientMap) Delete(key int64) {
	clientMap.m.Delete(key)
}

func (clientMap *ClientMap) Load(key int64) (value *Node, ok bool) {
	v, ok := clientMap.m.Load(key)
	if v != nil {
		value = v.(*Node)
	}
	return
}

func (clientMap *ClientMap) LoadOrStore(key int64, value *Node) (actual string, loaded bool) {
	a, loaded := clientMap.m.LoadOrStore(key, value)
	actual = a.(string)
	return
}

func (clientMap *ClientMap) Range(f func(key int64, value *Node) bool) {
	f1 := func(key, value interface{}) bool {
		return f(key.(int64), value.(*Node))
	}
	clientMap.m.Range(f1)
}

func (clientMap *ClientMap) Store(key int64, value *Node) {
	clientMap.m.Store(key, value)
}
