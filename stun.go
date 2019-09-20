package turn

import (
	"net"

	"github.com/hcm007/turn/internal/ipnet"
	"github.com/pion/stun"
)

// caller must hold the mutex
func (s *Server) handleBindingRequest(conn net.PacketConn, srcAddr net.Addr, m *stun.Message) error {
	s.log.Debugf("received BindingRequest from %s", srcAddr.String())
	ip, port, err := ipnet.AddrIPPort(srcAddr)
	if err != nil {
		return err
	}

	attrs := s.makeAttrs(m.TransactionID, stun.BindingSuccess, &stun.XORMappedAddress{
		IP:   ip,
		Port: port,
	}, stun.Fingerprint)

	return s.sender(conn, srcAddr, attrs...)
}
