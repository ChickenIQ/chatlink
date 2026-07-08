package socket

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/chickeniq/chatlink/pkg/proto"
)

func NewSocket(socketPath string, handler SockerHandler) (*Socket, error) {
	if handler == nil {
		return nil, fmt.Errorf("event handler cannot be nil")
	}

	if err := os.Remove(socketPath); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to remove existing socket file: %w", err)
	}

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		return nil, err
	}

	s := &Socket{
		listener: listener,
		handler:  handler,
		conn:     nil,
	}

	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				return
			}

			s.setConn(conn)
			s.acceptConn(conn)
		}
	}()

	return s, nil
}

func (s *Socket) acceptConn(conn net.Conn) {
	defer s.closeConn()

	for {
		p := proto.Packet{}
		if err := binary.Read(conn, binary.LittleEndian, &p); err != nil {
			log.Println("Error while reading from the socket", err)
			return
		}

		if int(p.Len) > proto.MaxPayloadSize {
			log.Printf("Paylod too big")
			return
		}

		if err := s.handlePacket(&p); err != nil {
			log.Println(err)
			return
		}
	}
}

func (s *Socket) closeConn() {
	s.Lock()
	defer s.Unlock()
	if s.conn == nil {
		return
	}

	s.conn.Close()
	s.conn = nil
}

func (s *Socket) setConn(conn net.Conn) {
	s.Lock()
	defer s.Unlock()
	s.conn = conn
}

func (s *Socket) Close() error {
	s.closeConn()
	return s.listener.Close()
}
