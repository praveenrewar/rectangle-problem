package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/suite"
)

type StdOutSuite struct {
	suite.Suite
	so osFMT
}

func (s *StdOutSuite) SetupSuite() {
	s.so = osFMT{}
}

func TestStdOutSuite(t *testing.T) {
	s := new(StdOutSuite)
	suite.Run(t, s)
}

func (s *StdOutSuite) TestPrintRectangles() {
	var rectangles []Rectangle
	rectangles = append(rectangles, Rectangle{X: 100, Y: 100, W: 250, H: 80})
	buf := &bytes.Buffer{}
	s.so.out = buf
	s.so.PrintRectangles(rectangles)
	actual := buf.String()
	if strings.Contains(actual, "1: Rectangle at (100, 100), w=250, h=80.") == false {
		s.T().Errorf("Expected: %v got: %v", "w=250, h=80.", actual)
	}
}

func (s *StdOutSuite) TestPrintIntersections() {
	var intersections []Intersection
	intersections = append(intersections, Intersection{RectangleNumbers: []int{0, 1, 2}, Rectangle: Rectangle{X: 100, Y: 100, W: 250, H: 80}})
	buf := &bytes.Buffer{}
	s.so.out = buf
	s.so.PrintIntersections(intersections)
	actual := buf.String()
	if strings.Contains(actual, "1: Between rectangle 1, 2 and 3 at (100, 100), w=250, h=80.") == false {
		s.T().Errorf("Expected: %v got: %v", "w=250, h=80.", actual)
	}

}
