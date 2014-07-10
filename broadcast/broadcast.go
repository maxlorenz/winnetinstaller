package broadcast

import (
	"fmt"
	"net"
	"strings"
)

func Broadcast(address string) {
	addr, err := net.ResolveUDPAddr("udp", address)
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}

	fmt.Fprintf(conn, "%v", strings.Split(conn.LocalAddr().String(), ":")[0])
}

func IpToBroadcast(ip string) string {
	ipSplit := strings.Split(ip, ".")
	return strings.Join(ipSplit[0:3], ".") + ".255"
}
