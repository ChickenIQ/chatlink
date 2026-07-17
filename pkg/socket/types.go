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
	HandleBotInfo(bots []proto.BotInfo)
	HandleKicked(reason proto.Kicked)
	HandleMessage(msg proto.Message)
}
