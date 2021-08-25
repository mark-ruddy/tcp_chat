package server

import (
  "fmt"
)

func (server *Server) doesRoomExist(checkRoom string) bool {
  for _, room := range server.rooms {
    if room == checkRoom {
      return true
    }
  }
  return false
}

func (server *Server) addClientToRoom(client *Client, desiredRoom string) {
  exists := server.doesRoomExist(desiredRoom)
  if !exists {
    server.sendClient(client, fmt.Sprintf("The room you requesed room to join '%s' does not exist", desiredRoom))
    return
  }
  
  // Is the user already in this room?
  if client.room == desiredRoom {
    server.sendClient(client, fmt.Sprintf("You are already in room %s", desiredRoom))
    return
  }

  client.room = desiredRoom
  server.sendClient(client, fmt.Sprintf("You have been added to room %s", desiredRoom))
}

func (server *Server) createRoom(client *Client, desiredRoom string) {
  exists := server.doesRoomExist(desiredRoom)
  if exists {
    server.sendClient(client, fmt.Sprintf("The room you requested to create '%s' already exists", desiredRoom))
    return
  }

  server.rooms = append(server.rooms, desiredRoom)
  server.sendClient(client, fmt.Sprintf("Room '%s' has been created - you can now enter using: /join %s", desiredRoom, desiredRoom))
}

