package socket

import (
	"net"
	"sync"
	"time"

	"github.com/chickeniq/chatlink/pkg/proto"
)

type Socket struct {
	timeout  time.Duration
	listener net.Listener
	mu       sync.Mutex
	conn     net.Conn
	path     string

	signedMessageHandlers []func(proto.SignedMessage)
	disconnectHandlers    []func(proto.Disconnect)
	botInfoHandlers       []func([]proto.BotInfo)
	messageHandlers       []func(proto.Message)
}
