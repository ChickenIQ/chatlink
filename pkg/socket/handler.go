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

	case proto.PacketKicked:
		s.handler.HandleKicked(proto.Kicked{
			ID:      p.Id,
			Content: string(p.Payload[:p.Len]),
		})

	case proto.PacketBotInfo:
		if int(p.Len)%proto.MaxUsernameSize != 0 {
			return fmt.Errorf("invalid bot info payload length: %d", p.Len)
		}

		count := int(p.Len) / proto.MaxUsernameSize
		if count > 127 {
			return fmt.Errorf("too many usernames: %d", count)
		}

		bots := make([]proto.BotInfo, count)
		for i := range bots {
			start := i * proto.MaxUsernameSize
			username := p.Payload[start : start+proto.MaxUsernameSize]

			bots[i] = proto.BotInfo{
				ID:      int8(i + 1),
				Content: strings.TrimRight(string(username), "\x00"),
			}
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
