package common

import (
	"sync"

	"github.com/gorilla/websocket"
)

// 构造连接
type Node struct {
	Conn      *websocket.Conn //socket连接
	Addr      string          //客户端地址
	DataQueue chan []byte     //消息内容
	// GroupSets set.Interface   //好友 / 群

}

var clientMap map[int64]*Node = make(map[int64]*Node, 0)

// 读写锁，绑定node时需要线程安全
var rwLocker sync.RWMutex
