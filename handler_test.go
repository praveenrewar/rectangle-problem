package main

import "testing"

func TestGetIntersections(t *testing.T) {
	mockFS := InitMockFileSystem()
	mockStdOut := InitMockStdOut()
	var input = `{"rects": [{"x": 100, "y": 100, "w": 250, "h": 80}, {"x": 120, "y": 200, "w": 250, "h": 150}]}`
	mockFS.On("ReadFile", InputFile).Return([]byte(input), nil)
	var rectangles []Rectangle
	rectangles = append(rectangles, Rectangle{X: 100, Y: 100, W: 250, H: 80})
	rectangles = append(rectangles, Rectangle{X: 120, Y: 200, W: 250, H: 150})
	var intersections []Intersection
	intersections = append(intersections, Intersection{RectangleNumbers: []int{0, 1}, Rectangle: Rectangle{X: 120, Y: 100, W: 230, H: 50}})
	mockStdOut.On("PrintRectangles", rectangles)
	mockStdOut.On("PrintIntersections", intersections)
	getIntersections()
	mockFS.AssertExpectations(t)
	mockStdOut.AssertExpectations(t)
}
