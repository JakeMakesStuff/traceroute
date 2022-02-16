package traceroute

import (
	"net"
	"syscall"
	"testing"
)

func TestAddrToSockAddr(t *testing.T) {
	tests := []struct {
		name string

		ip net.IP
	}{
		{
			name: "ipv4",
			ip:   net.ParseIP("173.194.72.99"),
		},
		{
			name: "ipv6",
			ip:   net.ParseIP("2606:4700:4700::1001"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			origIP := tt.ip
			origPort := 34456
			s := ipPortToSockaddr(origIP, origPort)
		
			var ip net.IP
			switch t := s.(type) {
			case *syscall.SockaddrInet4:
				ip = net.IP(t.Addr[:])
			case *syscall.SockaddrInet6:
				ip = net.IP(t.Addr[:])
			}
			if !ip.Equal(origIP) {
				t.Errorf("IPs don't match! orig=%v parsed=%v", origIP.String(), ip.String())
			}
		})
	}
}
