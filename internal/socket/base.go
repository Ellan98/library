package socket

import (
	"github.com/gorilla/websocket"
)

type SocketContrll struct {
	// enter  string //进入
	// exit   string //退出
	// pause  string //暂停
	// resume string //恢复
	Conn      *websocket.Conn //socket连接
	Addr      string          //客户端地址
	DataQueue chan []byte     //消息内容
	// GroupSets set.Interface   //好友 / 群

}

func (socket *SocketContrll) Connect() {

}
