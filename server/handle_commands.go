package server

import (
  "fmt"
  "strings"
)

func (server *Server) sendNoPermissionMsg(client *Client) {
  server.sendClient(client, "You do not have admin privileges to execute that command")
}

func (server *Server) handleClientMsg(client *Client, msg string) error {
  if len(msg) == 0 {
    return fmt.Errorf("empty string sent to handleClientMsg")
  }
  if msg[0] != '/' {
    server.broadcastToRoom(fmt.Sprintf("%s: %s", client.username, msg), client.room)
    return nil
  }

  // with a / prefix we know it is a command request
  // remove / now for processing
  msg = msg[1:]
  msgSplit := strings.Split(msg, " ")
  cmd := msgSplit[0]

  firstArg := ""
  if len(msgSplit) > 1 {
    firstArg = msgSplit[1]
  }

  switch cmd {
  case "admin":
    // TODO: Not sure how easy it would be too spoof a localhost ipAddr, this is likely insecure
    if client.ipAddr == "127.0.0.1" {
      client.Admin = true
      server.sendClient(client, "You have assumed admin privileges")
    } else {
      server.sendClient(client, "You do not have permission to assume admin privileges")
    }
  case "destroy":
    if client.Admin && firstArg != "" {
      server.destroyClientFromUUID(firstArg, client)
    } else if client.Admin && firstArg == "" {
      server.sendClient(client, "Usage: /destroy <UUID>")
    } else {
      server.sendNoPermissionMsg(client)
    }
  case "join":
    if firstArg != "" {
      server.addClientToRoom(client, firstArg)
    } else {
      server.sendClient(client, "Usage: /join <room_name>")
    }
  case "create":
    if firstArg != "" {
      server.createRoom(client, firstArg)
    } else {
      server.sendClient(client, "Usage: /create <room_name>")
    }
  case "list":
    server.sendClientList(client)
  case "rooms":
    server.sendClientRooms(client)
  case "help":
    server.sendClient(client, "No help docs available right now")
  default:
    server.sendClient(client, fmt.Sprintf("Command %s does not exist", msg))
  }
  return nil
}

