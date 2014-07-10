package main

import (
  "net"
  "fmt"
)

func main() {
  addr := net.UDPAddr {
    Port: 12345,
    IP: net.ParseIP("10.43.64.213"),
  }

  conn, err := net.ListenUDP("udp", &addr)
  defer conn.Close()
  if err != nil {
    panic(err)
  }

  for {
    buf := make([]byte, 1024)

    res, err := conn.Read(buf)

    if err != nil {
      panic(err)
    }

    if res != 0 {
      fmt.Printf("%v\n", string(buf[:res]))
    }
  }

}