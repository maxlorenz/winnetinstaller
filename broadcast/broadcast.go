package broadcast

import (
  "net"
  "fmt"
  "strings"
)

func BroadcastIP(address string) {
  addr, err := net.ResolveUDPAddr("udp", address)
  conn, err := net.DialUDP("udp", nil, addr)
  if err != nil {
    panic(err)
  }

  fmt.Fprintf(conn, "%v", strings.Split(conn.LocalAddr().String(), ":")[0])
}