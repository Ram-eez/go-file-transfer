package models

import (
	"encoding/json"
	"log"
	"os"
)

type File struct {
	Location string
	Name     string
	Size     string
}

func (f *File) OpenFile(location string, name string) *os.File {
	data, err := os.Open(location + "/" + name)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func (f *File) Serilalize() ([]byte, error) {
	data, err := json.Marshal(f)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return data, nil
}

func (f *File) CreateFile() (*os.File, error) {
	file, err := os.Create("/home/rameez/Downloads/" + f.Name)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return file, nil
}

