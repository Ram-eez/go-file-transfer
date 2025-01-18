package client

import (
	"fmt"
	"go-file-transfer/models"
	"io"
	"log"
	"net"
	"sync"
)

func TCPUpload(file *models.File, wg *sync.WaitGroup, semaphore chan struct{}, conn net.Conn) {

	defer wg.Done()
	semaphore <- struct{}{}

	metadata, err := file.Serilalize()
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write(metadata)
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)
	data := file.OpenFile(file.Location, file.Name)

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
