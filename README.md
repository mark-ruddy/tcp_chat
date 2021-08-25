# tcp_chat

Chatroom server designed to allow room creation/switching for connected users. 

### Usage

If no port number is provided, server will default to using 9999.

Build server and run binary on linux:
```
go build -o tcp_chat main.go
./tcp_chat <PORT>
```

Connect to server using netcat/telnet:
```
netcat <HOST> <PORT>
netcat localhost 9999
```

New users will be defaulted to room 'main', and can create new rooms at will:
```
/create myNewRoom
Room 'myNewRoom' has been created - you can now enter using: /join myNewRoom
/join myNewRoom
You have been added to room myNewRoom
```

### Administration

Users connected by localhost(127.0.0.1) IP can request admin privileges with `/admin`, allowing them to manage the server, remove other users etc.
