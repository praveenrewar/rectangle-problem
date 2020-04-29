package main

import (
	"io/ioutil"
)

var fs fileSystem

type fileSystem interface {
	ReadFile(name string) ([]byte, error)
}

// osFS implements fileSystem using the local disk.
type osFS struct{}

func (fs osFS) ReadFile(name string) ([]byte, error) {
	bytes, err := ioutil.ReadFile(name)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

//InitFileSystem is used to initialise the fileSystem
func InitFileSystem(FS fileSystem) {
	fs = FS
}
