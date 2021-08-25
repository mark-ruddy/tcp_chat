# tcp_chat

Currently a WIP

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

### Administration

Users connected by localhost(127.0.0.1) IP can request admin privileges with `/admin`, allowing them to manage the server, remove over users etc.
