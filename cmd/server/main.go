package main

import (
	"flag"
	"log"
	"net"
	"strconv"
)

func main() {
	host := flag.String("host", "127.0.0.1", "host IP. Loopback range is 127.x.y.z")
	port := flag.Int("port", 8080, "listening port")
	flag.Parse()

	l, err := net.Listen("tcp", net.JoinHostPort(*host, strconv.Itoa(*port)))
	if err != nil {
		log.Fatal("listening:", err)
	}
	defer l.Close()

	log.Println("server listing on", l.Addr().String())
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("accepting: ", err)
		}
		go onNewConnection(conn)
	}
}

func onNewConnection(conn net.Conn) {
	defer conn.Close()
	log.Println("server connection", conn.RemoteAddr().String(), "=>", conn.LocalAddr().String())
}
