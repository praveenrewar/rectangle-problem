package main

import (
	"fmt"
	"io"
)

type StdOut interface {
	PrintRectangles(rectangles []Rectangle)
	PrintIntersections(intersections []Intersection)
}

var so StdOut

type osFMT struct {
	out io.Writer
}

func (so osFMT) PrintRectangles(rectangles []Rectangle) {
	fmt.Fprintln(so.out, "Input:")
	for i, rectangle := range rectangles {
		fmt.Fprintf(so.out, "%v: Rectangle at (%v, %v), w=%v, h=%v.\n", i+1, rectangle.X, rectangle.Y, rectangle.W, rectangle.H)
	}
}

func (so osFMT) PrintIntersections(intersections []Intersection) {
	fmt.Fprintln(so.out, "Intersections:")
	for i, intersection := range intersections {
		var intersectingRectangles string
		intersectingRectangles = fmt.Sprintf("%v", intersection.RectangleNumbers[0]+1)
		for i := 1; i < len(intersection.RectangleNumbers)-1; i++ {
			intersectingRectangles += fmt.Sprintf(", %v", intersection.RectangleNumbers[i]+1)
		}
		intersectingRectangles += fmt.Sprintf(" and %v", intersection.RectangleNumbers[len(intersection.RectangleNumbers)-1]+1)
		fmt.Fprintf(so.out, "%v: Between rectangle %v at (%v, %v), w=%v, h=%v.\n", i+1, intersectingRectangles, intersection.Rectangle.X, intersection.Rectangle.Y, intersection.Rectangle.W, intersection.Rectangle.H)
	}
}

func InitStdOut(SO StdOut) {
	so = SO
}
