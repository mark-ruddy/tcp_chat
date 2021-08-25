package server

import (
  "fmt"
  "net"
  "log"
  "github.com/nu7hatch/gouuid"
  "github.com/TwinProduction/go-color"
)

type Client struct {
  id *uuid.UUID
  currentRoom *uuid.UUID
  ipAddr string
  username string
  conn net.Conn

  Admin bool
}

type Server struct {
  port string
  clients map[*uuid.UUID]Client
  clientCount int
  currentClient *uuid.UUID
  mainRoom *uuid.UUID
  rooms []*uuid.UUID

  Delim string
}

func NewServer(port string) Server {
  mainRoom, err := uuid.NewV4()
  if err != nil {
    log.Printf("error creating mainRoom UUID: %s", err)
  }

  return Server {
    port,
    make(map[*uuid.UUID]Client),
    0,
    nil,
    mainRoom,
    []*uuid.UUID{mainRoom},
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

    client := Client{ UUID, server.mainRoom, conn.RemoteAddr().String(), "", conn, false }
    fmt.Println(color.Ize(color.Green, "New connection from " + client.ipAddr))
    fmt.Println(color.Ize(color.Green, "Client created with UUID: " + client.id.String()))

    server.clients[UUID] = client 
    go server.handleClient(&client)
  }
}

func (server *Server) handleClient(client *Client) {
  server.sendClient(*client, "Welcome! Enter your username")
  client.username = server.recvClient(*client)
  server.broadcast(fmt.Sprintf("%s has joined the chat", client.username))

  for {
    msg := server.recvClient(*client)
    err := server.handleClientMsg(client, msg)
    if err != nil {
      log.Printf("error in handleClientMsg triggered: %s", err)
      server.destroyClient(*client)
      break
    }
  }
}

