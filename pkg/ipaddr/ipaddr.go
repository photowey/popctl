package ipaddr

import (
	"fmt"
	"net"
)

var LocalIp = ""

func init() {
	LocalIp, _ = ParseLocalIP()
}

func ParseLocalIP() (string, error) {
	ip := ""
	interfaces, err := net.Interfaces()
	if err != nil {
		return "", fmt.Errorf("acquire net interfaces error: %v", err)
	}

LABEL:
	for _, intf := range interfaces {
		if ((intf.Flags & net.FlagUp) != 0) && ((intf.Flags & net.FlagLoopback) == 0) {
			addrs, err := intf.Addrs()
			if err != nil {
				return "", fmt.Errorf("acquired net interface.addr error: %v", err)
			}
			for _, address := range addrs {
				if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
					ip = ipNet.IP.String()
					break LABEL
				}
			}
		}
	}

	return ip, nil
}
