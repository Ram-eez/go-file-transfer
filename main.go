package main

import (
	"go-file-transfer/client"
	"go-file-transfer/server"
	"time"
)

func main() {
	go server.TCPListenAndAccept()

	time.Sleep(1 * time.Second)

	client.TCPDail()

	time.Sleep(1 * time.Second)

}
