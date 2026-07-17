package proto

import (
	"fmt"
)

func NewPacket(packetType PacketType, id int8, payload []byte) (*Packet, error) {
	p := &Packet{
		PacketType: packetType,
		ID:         id,
		Len:        uint16(len(payload)),
	}

	if len(payload) > MaxPayloadSize {
		return nil, fmt.Errorf("payload size exceeds maximum allowed size")
	}

	copy(p.Payload[:], payload)

	return p, nil
}

func MessagePacket(id int8, content string) (*Packet, error) {
	return NewPacket(PacketMessage, id, []byte(content))
}

func DisconnectPacket() *Packet {
	return &Packet{PacketType: PacketDisconnect}
}

func StopPacket() *Packet {
	return &Packet{PacketType: PacketStop}
}
