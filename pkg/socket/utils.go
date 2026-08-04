package socket

import (
	"net"
)

func (s *Socket) setConn(conn net.Conn) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.listener != nil {
		s.conn = conn
		return nil
	}

	if conn != nil {
		conn.Close()
	}

	return net.ErrClosed
}

func (s *Socket) closeConn(conn net.Conn) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.conn == conn {
		s.conn = nil
	}

	if conn != nil {
		conn.Close()
	}
}

func (s *Socket) Close() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.listener != nil {
		s.listener.Close()
		s.listener = nil
	}

	if s.conn != nil {
		s.conn.Close()
		s.conn = nil
	}
}
