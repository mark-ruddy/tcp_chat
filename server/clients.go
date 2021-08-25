package server

import (
  "fmt"
  "bufio"
  "log"
  "strings"
  "github.com/TwinProduction/go-color"
  "github.com/nu7hatch/gouuid"
)

func (server *Server) recvClient(client Client) string {
  incData, err := bufio.NewReader(client.conn).ReadString([]byte(server.Delim)[0])
  if err != nil {
    log.Printf("recvClient bufio.NewReader creation failed: %s", err)
  }

  incMsg := strings.TrimSpace(string(incData))
  if len(incMsg) == 0 {
    return ""
  }

  if server.Delim == "\n" {
    // \n whitespace already TrimSpace'd off
    return incMsg
  } else {
    // else remove delim manually
    return incMsg[:len(incMsg) - len(server.Delim)]
  }
}

func (server *Server) sendClient(client Client, msg string) {
  client.conn.Write([]byte(msg + server.Delim))
}

func (server *Server) destroyClient(client Client) {
  fmt.Println(color.Ize(color.Red, "Client " + client.id.String() + " removed"))
  server.sendClient(client, "destroy")
  delete(server.clients, client.id)
}

func (server *Server) destroyClientFromUUID(UUID string) {
  actualUUID, err := uuid.Parse([]byte(UUID))
  if err != nil {
    log.Printf("error parsing UUID in destroyClientFromUUID: %s", err)
  }

  server.destroyClient(server.clients[actualUUID])
}

func (server *Server) broadcast(msg string) {
  for _, client := range server.clients {
    server.sendClient(client, msg)
  }
}

