package client

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"sync"
)

func TCPUpload(filepath string, wg *sync.WaitGroup, semaphore chan struct{}, conn net.Conn) {

	defer wg.Done()
	semaphore <- struct{}{}

	buf := make([]byte, 1024)
	data := OpenFile(filepath)

	for {
		n, err := data.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if err == io.EOF {
			break
		}

		_, err = conn.Write(buf[:n])
		if err != nil {
			log.Fatal(err)
		}

	}
	<-semaphore
	fmt.Println("Data sent to server")
}

func OpenFile(filepath string) *os.File {
	data, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
