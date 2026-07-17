package socket

import (
	"net"
	"sync"

	"github.com/chickeniq/chatlink/pkg/proto"
)

type Socket struct {
	handler  SockerHandler
	listener net.Listener
	conn     net.Conn
	sync.Mutex
}

type SockerHandler interface {
	HandleSignedMessage(msg proto.SignedMessage)
	HandleDisconnect(reason proto.Disconnect)
	HandleBotInfo(bots []proto.BotInfo)
	HandleMessage(msg proto.Message)
}
