package main

import "github.com/stretchr/testify/mock"

//MockFileSystem is a mock file system
type MockFileSystem struct {
	mock.Mock
}

//ReadFile is mock function for ReadFile
func (m *MockFileSystem) ReadFile(name string) ([]byte, error) {
	rets := m.Called(name)
	return rets.Get(0).([]byte), rets.Error(1)
}

//InitMockFileSystem initializes the fs variable, it assigns
//a new MockFileSystem instance to it, instead of an actual fileSystem
func InitMockFileSystem() *MockFileSystem {
	s := new(MockFileSystem)
	fs = s
	return s
}
