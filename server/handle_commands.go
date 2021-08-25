package server

import (
  "fmt"
  "strings"
)

func (server *Server) sendNoPermissionMsg(client Client) {
  server.sendClient(client, "You do not have admin privileges to execute that command")
}

func (server *Server) handleClientMsg(client *Client, msg string) error {
  if len(msg) == 0 {
    return fmt.Errorf("empty string sent to handleClientMsg")
  }
  if msg[0] != '/' {
    server.broadcast(fmt.Sprintf("%s: %s", client.username, msg))
    return nil
  }

  // with a / prefix we know it is a command request
  // remove / now for processing
  msg = msg[1:]
  msgSplit := strings.Split(msg, " ")
  cmd := msgSplit[0]
  firstArg := msgSplit[1]

  switch cmd {
  case "admin":
    // TODO: Not sure how easy it would be too spoof a localhost ipAddr, this is likely insecure
    if client.ipAddr == "127.0.0.1" {
      client.Admin = true
    } else {
      server.sendClient(*client, "You do not have permission to assume admin privileges")
    }
  case "destroy":
    if client.Admin {
      server.destroyClientFromUUID(firstArg)
    } else {
      server.sendNoPermissionMsg(*client)
    }
  case "help":
    server.sendClient(*client, "No help docs available right now")
  default:
    server.sendClient(*client, fmt.Sprintf("Command %s does not exist", msg))
  }
  return nil
}

