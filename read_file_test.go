package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type FileSuite struct {
	suite.Suite
	file osFS
}

func (f *FileSuite) SetupSuite() {
	f.file = osFS{}
}

func TestFileSuite(t *testing.T) {
	s := new(FileSuite)
	suite.Run(t, s)
}

func (f *FileSuite) TestReadFile() {
	filebytes, err := f.file.ReadFile("test.txt")
	if err != nil {
		f.T().Fatal(err)
	}
	if string(filebytes) != "This is test file for testing ReadFile function" {
		f.T().Errorf("Error in reading file")
	}

	//Test ReadFile failure
	filebytes, err = f.file.ReadFile("tes.txt")
	if err == nil {
		f.T().Errorf("Expected error in reading file")
	}
}
