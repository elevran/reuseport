package main

import (
	"flag"
	"log"
	"time"

	reuse "github.com/libp2p/go-reuseport"
)

func main() {
	client := flag.String("client", "127.0.0.1:0", "client bind IP and port")
	server := flag.String("server", "127.0.0.2:8080", "server IP and port")
	flag.Parse()

	conn, err := reuse.Dial("tcp", *client, *server)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()
	log.Println("client connection", conn.RemoteAddr().String(), "=>", conn.LocalAddr().String())
	time.Sleep(1 * time.Minute)
}
