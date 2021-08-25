package server

import (
  "fmt"
  "net"
  "log"
  "strings"
  "github.com/nu7hatch/gouuid"
  "github.com/TwinProduction/go-color"
)

type Client struct {
  UUID *uuid.UUID
  room string
  ipAddr string
  username string
  conn net.Conn

  Admin bool
}

type Server struct {
  port string
  clients map[*uuid.UUID]*Client
  clientCount int
  currentClient *uuid.UUID
  mainRoom string
  rooms []string

  Delim string
}

func NewServer(port string) Server {
  mainRoom := "main"

  return Server {
    port,
    make(map[*uuid.UUID]*Client),
    0,
    nil,
    mainRoom,
    []string{mainRoom},
    "\n",
  }
}

func (server *Server) Listen() {
  s, _ := net.Listen("tcp", server.port)
  fmt.Println("Server listening on port ", server.port)

  for {
    conn, err := s.Accept()
    if err != nil {
      log.Printf("error accepting client: %s", err)
    }
    server.clientCount++
    UUID, err := uuid.NewV4()
    if err != nil {
      log.Printf("error creating client UUID: %s", err)
    }

    client := Client{ UUID, server.mainRoom, strings.Split(conn.RemoteAddr().String(), ":")[0], "", conn, false }
    fmt.Println(color.Ize(color.Green, "New connection from " + client.ipAddr))
    fmt.Println(color.Ize(color.Green, "Client created with UUID: " + client.UUID.String()))

    server.clients[UUID] = &client 
    go server.handleClient(server.clients[UUID])
  }
}

func (server *Server) handleClient(client *Client) {
  server.sendClient(client, "Welcome! Enter your username")
  client.username = server.recvClient(client)
  server.broadcast(fmt.Sprintf("%s has connected joined chatroom: %s", client.username, client.room))

  for {
    msg := server.recvClient(client)
    err := server.handleClientMsg(client, msg)
    if err != nil {
      log.Printf("error in handleClientMsg triggered: %s", err)
      server.destroyClient(client)
      break
    }
  }
}

