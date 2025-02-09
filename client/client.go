package client

import (
	"go-file-transfer/models"
	"log"
	"net"
	"os"
	"sync"
)

func TCPDail() {
	addr, err := net.ResolveTCPAddr("tcp", ":3000")
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

	fileLocation := "/home/rameez/Downloads/New Folder/DespicableMe/"
	fileName := "DespicableMe.mp4"
	// file.FileSet(fileLocation, fileName)

	file := FileInit(fileLocation, fileName)

	models.Files = append(models.Files, *file)

	for _, file := range models.Files {
		wg.Add(1)
		go TCPUpload(&file, &wg, semaphore, conn)
	}

	wg.Wait()
}

func FileInit(FilePath string, FileName string) *models.File {

	FileSize, err := os.Stat(FilePath + "/" + FileName)
	if err != nil {
		log.Fatal(err)
	}

	file := &models.File{
		Location: FilePath,
		Name:     FileName,
		Size:     FileSize.Size(),
	}

	return file
}
