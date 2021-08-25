package server

import (
  "fmt"
  "bufio"
  "log"
  "strings"
  "github.com/TwinProduction/go-color"
  "github.com/nu7hatch/gouuid"
)

func (server *Server) recvClient(client *Client) string {
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

func (server *Server) sendClient(client *Client, msg string) {
  client.conn.Write([]byte(msg + server.Delim))
}

func (server *Server) destroyClient(client *Client) {
  fmt.Println(color.Ize(color.Red, "Client " + client.UUID.String() + " removed"))
  server.sendClient(client, "You have been removed from the server")
  delete(server.clients, client.UUID)
}

func (server *Server) destroyClientFromUUID(UUID string, requester *Client) {
  actualUUID, err := uuid.Parse([]byte(UUID))
  if err != nil {
    log.Printf("error parsing UUID in destroyClientFromUUID: %s", err)
    server.sendClient(requester, "The UUID passed was invalid")
    return
  }
  client := server.clients[actualUUID]
  server.sendClient(requester, fmt.Sprintf("REMOVING: Username %s - UUID %s", client.username, UUID))
  server.destroyClient(client)
}

func (server *Server) broadcast(msg string) {
  for _, client := range server.clients {
    server.sendClient(client, fmt.Sprintf("<%s> %s", "broadcast", msg))
  }
}

func (server *Server) broadcastToRoom(msg string, room string ) {
  for _, client := range server.clients {
    if client.room == room {
      server.sendClient(client, fmt.Sprintf("<%s> %s", room, msg))
    }
  }
}

func (server *Server) sendClientList(requester *Client) {
  server.sendClient(requester, fmt.Sprintf("There are %d connected users", len(server.clients)))

  // Only admin clients should be able to view UUIDs of other clients
  if requester.Admin == true {
    for _, client := range server.clients {
      server.sendClient(requester, fmt.Sprintf("%s(UUID: %s) - room: %s", 
        client.username, 
        client.UUID.String(), 
        client.room,
      ))
    }
  } else {
    for _, client := range server.clients {
      server.sendClient(requester, fmt.Sprintf("%s is in room %s", 
        client.username, 
        client.room,
      ))
    }
  }
}

func (server *Server) sendClientRooms(client *Client) {
  server.sendClient(client, fmt.Sprintf("There are %d available rooms", len(server.rooms)))
  server.sendClient(client, fmt.Sprintf("You are currently in room: %s", client.room))
  server.sendClient(client, "\nAll rooms listed below: ")

  for _, room := range server.rooms {
    server.sendClient(client, fmt.Sprintf("- %s", room))
  }
}

