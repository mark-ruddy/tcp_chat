package main

import (
  "os"
  "tcp_chat/server"
)

func main() {
  port := ":9999"
  if len(os.Args[1:]) >= 1 {
    port = ":" + os.Args[1]
  }

  s := server.NewServer(port)
  s.Listen()
}

