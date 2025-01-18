package server

import (
	"encoding/json"
	"fmt"
	"go-file-transfer/models"
	"io"
	"log"
	"net"
)

func HandleConn(conn net.Conn) {

	defer conn.Close()
	metadata := make([]byte, 1024)

	n, err := conn.Read(metadata)
	if err != nil {
		log.Fatal(err)
	}

	var file models.File
	if err := json.Unmarshal(metadata[:n], &file); err != nil {
		log.Fatal(err)
	}

	emptyFile, err := file.CreateFile()
	if err != nil {
		log.Fatal(err)
	}

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		if err == io.EOF {
			bytesWritten, err := emptyFile.Write(buf[:n])
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%d bytes written\n", bytesWritten)
			break
		}

		bytesWritten, err := emptyFile.Write(buf[:n])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d bytes written\n", bytesWritten)

	}

}
