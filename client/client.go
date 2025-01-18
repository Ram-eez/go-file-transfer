package client

import (
	"go-file-transfer/models"
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

	// data := []byte("Hello, server")

	fileLocation := "/home/rameez/Documents"
	fileName := "testfile.txt"

	file := &models.File{
		Location: fileLocation,
		Name:     fileName,
	}
	// file.FileSet(fileLocation, fileName)

	messageLength := 1

	for i := 0; i < messageLength; i++ {
		wg.Add(1)
		go TCPUpload(file, &wg, semaphore, conn)
	}

	wg.Wait()
}
