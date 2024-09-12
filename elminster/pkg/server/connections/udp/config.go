package udp

import "net"

type Server struct {
	IPv4 string
	Port int
}

func CreateServer(config Server) *net.UDPConn {
	addr := &net.UDPAddr{
		IP:   net.IP(config.IPv4),
		Port: config.Port,
		Zone: "",
	}

	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		panic(err)
	}

	return conn
}
