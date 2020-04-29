package main

import (
	"math"
)

//GetIntersectingRectangle is used to get the intersection between 2 rectangles
func GetIntersectingRectangle(rectangleA Rectangle, rectangleB Rectangle) (Rectangle, bool) {
	var rectangleC Rectangle
	x1, y1 := rectangleA.X, rectangleA.Y-rectangleA.H
	x2, y2 := rectangleA.X+rectangleA.W, rectangleA.Y
	x3, y3 := rectangleB.X, rectangleB.Y-rectangleB.H
	x4, y4 := rectangleB.X+rectangleB.W, rectangleB.Y
	x5 := math.Max(x1, x3)
	y5 := math.Max(y1, y3)
	x6 := math.Min(x2, x4)
	y6 := math.Min(y2, y4)
	if x5 >= x6 || y5 >= y6 {
		return rectangleC, false
	}
	// fmt.Println("(", x1, ", ", y1, ")")
	// fmt.Println("(", x2, ", ", y2, ")")
	// fmt.Println("(", x3, ", ", y3, ")")
	// fmt.Println("(", x4, ", ", y4, ")")
	// fmt.Println("(", x5, ", ", y5, ")")
	// fmt.Println("(", x6, ", ", y6, ")")
	rectangleC.X = x5
	rectangleC.Y = y6
	rectangleC.H = y6 - y5
	rectangleC.W = x6 - x5
	return rectangleC, true
}
