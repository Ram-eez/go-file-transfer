package models

import (
	"log"
	"os"
)

type File struct {
	Location string
	Name     string
	Size     string
}

// func (f File) FileSet(FileLocation string, FileName string) string {
// 	f.Location = FileLocation
// 	f.Name = FileName
// 	return f.Location + "/" + f.Name
// }

func (f *File) OpenFile(location string, name string) *os.File {
	data, err := os.Open(location + "/" + name)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
