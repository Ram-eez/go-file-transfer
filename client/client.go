package client

import (
	"log"
	"net"
	"sync"
)

func TCPDail() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:3000")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	var wg sync.WaitGroup
	concurrentLimit := 5
	semaphore := make(chan struct{}, concurrentLimit)
	data := make([]byte, 0)
	for {
		wg.Add(1)
		go TCPUpload(data, &wg, semaphore, conn)
	}
}
