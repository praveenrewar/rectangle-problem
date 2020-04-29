package main

import "github.com/stretchr/testify/mock"

type MockStdOut struct {
	mock.Mock
}

func (m *MockStdOut) PrintRectangles(rectangles []Rectangle) {
	_ = m.Called(rectangles)
	return
}

func (m *MockStdOut) PrintIntersections(intersections []Intersection) {
	_ = m.Called(intersections)
	return
}

func InitMockStdOut() *MockStdOut {
	s := new(MockStdOut)
	so = s
	return s
}
