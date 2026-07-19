package socket

import (
	"time"

	"github.com/chickeniq/chatlink/pkg/proto"
)

type Option func(*Socket)

func WithSignedMessageHandler(fn func(proto.SignedMessage)) Option {
	return func(s *Socket) {
		if fn != nil {
			s.signedMessageHandlers = append(s.signedMessageHandlers, fn)
		}
	}
}

func WithDisconnectHandler(fn func(proto.Disconnect)) Option {
	return func(s *Socket) {
		if fn != nil {
			s.disconnectHandlers = append(s.disconnectHandlers, fn)
		}
	}
}

func WithBotInfoHandler(fn func([]proto.BotInfo)) Option {
	return func(s *Socket) {
		if fn != nil {
			s.botInfoHandlers = append(s.botInfoHandlers, fn)
		}
	}
}

func WithMessageHandler(fn func(proto.Message)) Option {
	return func(s *Socket) {
		if fn != nil {
			s.messageHandlers = append(s.messageHandlers, fn)
		}
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(s *Socket) {
		s.timeout = timeout
	}
}
