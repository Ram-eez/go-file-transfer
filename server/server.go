package server

import (
	"fmt"
	"log"
	"net"
)

func TCPListenAndAccept() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	defer ln.Close()
	fmt.Println("Listening on port 3000")
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err) // Log error but don't stop the server
			continue                                        // Continue accepting other connections
		}

		go HandleConn(conn)
	}
}
