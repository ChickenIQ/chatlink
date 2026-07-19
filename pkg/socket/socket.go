package socket

import (
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/chickeniq/chatlink/pkg/proto"
)

func NewSocket(socketPath string, options ...Option) (*Socket, error) {
	if err := os.Remove(socketPath); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("failed to remove existing socket file: %w", err)
	}

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		return nil, err
	}

	sock := &Socket{
		timeout:  15 * time.Second,
		listener: listener,
	}

	for _, option := range options {
		if option != nil {
			option(sock)
		}
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}

			if err := sock.setConn(conn); err != nil {
				return
			}

			sock.acceptConn(conn)
		}
	}()

	return sock, nil
}

func (s *Socket) acceptConn(conn net.Conn) {
	defer s.closeConn(conn)

	for {
		p := proto.Packet{}
		if err := binary.Read(conn, binary.LittleEndian, &p); err != nil {
			log.Println("Error while reading from the socket", err)
			return
		}

		if err := s.handlePacket(&p); err != nil {
			log.Println(err)
			return
		}
	}
}
