package proto

const MaxPayloadSize = 32768
const MaxUsernameSize = 16

type PacketType uint8

const (
	PacketMessage PacketType = iota
	PacketBotInfo
	PacketStop
	PacketDisconnect
	PacketSignedMessage
)

type Packet struct {
	PacketType PacketType
	Id         int8
	Len        uint16
	Payload    [MaxPayloadSize]byte
}

type Output struct {
	ID      uint8
	Content string
}

type SignedMessage Output
type Message Output
type BotInfo Output
