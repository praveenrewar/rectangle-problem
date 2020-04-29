package main

//Rectangle defines a rectangle with x, y being one of the corners(top left) and w, h being width and height
type Rectangle struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	W float64 `json:"w"`
	H float64 `json:"h"`
}

//Rectangles contains a slice of Rectangle
type Rectangles struct {
	Rects []Rectangle `json:"rects"`
}

type Intersection struct {
	RectangleNumbers []int
	Rectangle        Rectangle
}
