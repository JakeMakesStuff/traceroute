package traceroute

import (
	"errors"
	"net"
	"syscall"
)

func GetLocalIP(v6 bool) (net.IP, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return nil, err
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && (ipnet.IP.To4() == nil) == v6 && !ipnet.IP.IsLoopback() {
			return ipnet.IP, nil
		}
	}
	return nil, errors.New("You do not appear to be connected to the Internet")
}

// TODO: change to AddrToSockaddr
func ipPortToSockaddr(ip net.IP, port int) syscall.Sockaddr {
	if ip4 := ip.To4(); ip4 != nil {
		var ip [4]byte
		copy(ip[:], ip4)
		return &syscall.SockaddrInet4{
			Port: port,
			Addr: ip,
		}
	} else if ip6 := ip.To16(); ip6 != nil {
		var ip [16]byte
		copy(ip[:], ip6)
		return &syscall.SockaddrInet6{
			Port: port,
			Addr: ip,
		}
	}
	return nil
}

func sockAddrToIP(s syscall.Sockaddr) net.IP {
	switch t := s.(type) {
	case *syscall.SockaddrInet4:
		return t.Addr[:]
	case *syscall.SockaddrInet6:
		return t.Addr[:]
	}
	return nil
}
