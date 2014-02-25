package services

import (
	"base/socket"
)

// 处理接受到的数据
func MessageHandler(channel socket.IChannel, protoPack *socket.ProtoPack) {

}

//客户端连接事件处理器
func ConnectedHandler(channel socket.IChannel) {

}

//客户端连接断开事件处理器
func DisconnectHandler(channel socket.IChannel) {

}
