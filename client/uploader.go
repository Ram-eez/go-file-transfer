package client

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func TCPUpload(data []byte, wg *sync.WaitGroup, semaphore chan struct{}, conn net.Conn) {
	defer wg.Done()
	semaphore <- struct{}{}
	_, err := conn.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	<-semaphore
	fmt.Println("Data sent to server")
}
