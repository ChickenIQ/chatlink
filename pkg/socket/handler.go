package socket

import (
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/chickeniq/chatlink/pkg/proto"
)

func (s *Socket) handlePacket(p *proto.Packet) error {
	switch p.PacketType {
	case proto.PacketMessage:
		s.handler.HandleMessage(proto.Message{
			ID:      p.Id,
			Content: string(p.Payload[:p.Len]),
		})

	case proto.PacketSignedMessage:
		s.handler.HandleSignedMessage(proto.SignedMessage{
			ID:      p.Id,
			Content: string(p.Payload[:p.Len]),
		})

	case proto.PacketBotInfo:
		if int(p.Len)%proto.MaxUsernameSize != 0 {
			return fmt.Errorf("invalid payload length")
		}

		if int(p.Len)/proto.MaxUsernameSize > 255 {
			return fmt.Errorf("too many usernames")
		}

		var id int8 = 1
		var bots []proto.BotInfo
		for i := 0; i < int(p.Len); i += proto.MaxUsernameSize {
			usernameBytes := p.Payload[i : i+proto.MaxUsernameSize]
			bots = append(bots, proto.BotInfo{
				ID:      id,
				Content: strings.TrimRight(string(usernameBytes), "\x00"),
			})
			id++
		}

		s.handler.HandleBotInfo(bots)

	default:
		return fmt.Errorf("unknown packet type: %d", p.PacketType)
	}

	return nil
}

func (s *Socket) SendPacket(message *proto.Packet) error {
	s.Lock()
	defer s.Unlock()

	if s.conn == nil {
		return fmt.Errorf("no active client connection")
	}

	if err := binary.Write(s.conn, binary.LittleEndian, message); err != nil {
		return fmt.Errorf("failed to send packet: %w", err)
	}

	return nil
}
