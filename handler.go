package main

import (
	"encoding/json"
	"fmt"
)

func getIntersections() {
	var fileContent Rectangles
	var rectangles []Rectangle
	var intersections []Intersection
	bytes, err := fs.ReadFile(InputFile)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(bytes, &fileContent)
	if err != nil {
		fmt.Println(err)
	}
	if len(fileContent.Rects) > 10 {
		rectangles = fileContent.Rects[:10]
	} else {
		rectangles = fileContent.Rects
	}
	for i := 0; i < len(rectangles); i++ {
		for j := i + 1; j < len(rectangles); j++ {
			rectangle, flag := GetIntersectingRectangle(rectangles[i], rectangles[j])
			if flag {
				intersections = append(intersections, Intersection{RectangleNumbers: []int{i, j}, Rectangle: rectangle})
			}
		}
	}
	var flag = true
	for i := 0; flag == true ;i++{
		if len(intersections) - 1 < i {
			flag = false
			continue
		}
		var j = intersections[i].RectangleNumbers[len(intersections[i].RectangleNumbers)-1]
		for k := j+1; k < len(rectangles); k++ {
			rectangle, flag := GetIntersectingRectangle(intersections[i].Rectangle, rectangles[k])
			if flag {
				numbers := append(intersections[i].RectangleNumbers, k)
				intersections = append(intersections, Intersection{RectangleNumbers: numbers, Rectangle: rectangle})
			}
		}
	}
	so.PrintRectangles(rectangles)
	so.PrintIntersections(intersections)

}
