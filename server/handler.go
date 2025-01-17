package server

import (
	"fmt"
	"log"
	"net"
)

func HandleConn(conn net.Conn) {
	data := make([]byte, 1024)
	message, err := conn.Read([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Message recived: %+v", message)
}
