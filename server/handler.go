package server

import (
	"encoding/json"
	"fmt"
	"go-file-transfer/models"
	"io"
	"log"
	"net"

	"github.com/schollz/progressbar/v3"
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

	// var TotalBytesSent int64

	bar := progressbar.NewOptions64(file.Size,
		// progressbar.OptionSetWriter(ansi.NewAnsiStdout()), //you should install "github.com/k0kubun/go-ansi"
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(true),
		progressbar.OptionSetWidth(15),
		progressbar.OptionSetDescription("[cyan][1/3][reset] Writing moshable file..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))

	for {
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}

		bytesWritten, err := emptyFile.Write(buf[:n])
		if err != nil {
			log.Fatal(err)
		}
		bar.Add(bytesWritten)

		if err == io.EOF {
			fmt.Printf("\nRecieved the data successfully")
			break
		}
	}

	bar.Finish()

}
