package main

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func TestGetIntersectingRectangles(t *testing.T) {
	rectangleA := Rectangle{
		X: 100,
		Y: 100,
		W: 250,
		H: 80,
	}
	rectangleB := Rectangle{
		X: 140,
		Y: 160,
		W: 250,
		H: 100,
	}
	rectangleC, flag := GetIntersectingRectangle(rectangleA, rectangleB)
	assert.Equal(t, flag, true)
	assert.Equal(t, rectangleC, Rectangle{X: 140, Y: 100, W: 210, H: 40})
}
