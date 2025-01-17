package server

import (
	"fmt"
	"log"
	"net"
)

func HandleConn(conn net.Conn) {

	defer conn.Close()

	data := make([]byte, 1024)
	messageLength, err := conn.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Message recived: %s", string(data[:messageLength]))
}
