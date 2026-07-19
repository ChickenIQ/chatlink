package socket

import (
	"fmt"
	"net"
)

func (s *Socket) Close() {
	s.closeListener()

	s.mu.Lock()
	conn := s.conn
	s.mu.Unlock()

	s.closeConn(conn)
}

func (s *Socket) closeConn(conn net.Conn) {
	s.mu.Lock()
	if s.conn == conn {
		s.conn = nil
	}
	s.mu.Unlock()

	if conn != nil {
		conn.Close()
	}
}

func (s *Socket) closeListener() {
	s.mu.Lock()
	listener := s.listener
	s.listener = nil
	s.mu.Unlock()

	if listener != nil {
		listener.Close()
	}
}

func (s *Socket) setConn(conn net.Conn) error {
	s.mu.Lock()
	if s.listener == nil {
		s.mu.Unlock()
		conn.Close()
		return fmt.Errorf("listener is nil")
	}

	s.conn = conn
	s.mu.Unlock()

	return nil
}
