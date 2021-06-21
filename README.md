# TCP REUSEPORT and REUSEADDR Options

Simple PoC for concurrent reuse of client IP and port.
 Since packets are associated with sockets using the the full tuple
 (e.g., source and destination IP:port), we can (safely?) connect multiple
 sockets to the same local IP:port as long as the destinations are different.
 When calling connect(), the kernel will pick a free ephemeral port, thus
 limiting the total number of outgoing connections to <64KB (16b port).
 Note that the same limitation does not apply to server sockets (i.e., a
 properly configured server can accept more than 64K concurrent connections).

To run the PoC:

```sh
# compile client and server
mkdir bin
go build -o bin/server ./cmd/server
go build -o bin/client ./cmd/client
# start multiple servers (using the 127.x.x.x loopback range)
./bin/server & # default ip:port (127.0.0.1:8080)
./bin/server -host 127.0.0.2 -port 8080 & # and a second server on a different loopback IP
# connect the client to both
./bin/client -client 127.0.0.1:1234 -server 127.0.0.1:8080 &
./bin/client -client 127.0.0.1:1234 -server 127.0.0.2:8080 &
# both clients and servers should output the local and remote connection addresses
# the clients sleep for 1m then close the connection and exist
```
