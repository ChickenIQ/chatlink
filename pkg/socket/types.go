package socket

import (
	"net"
	"sync"
	"time"

	"github.com/chickeniq/chatlink/pkg/proto"
)

type Socket struct {
	mu       sync.Mutex
	timeout  time.Duration
	listener net.Listener
	conn     net.Conn

	signedMessageHandlers []func(proto.SignedMessage)
	disconnectHandlers    []func(proto.Disconnect)
	botInfoHandlers       []func([]proto.BotInfo)
	messageHandlers       []func(proto.Message)
}
