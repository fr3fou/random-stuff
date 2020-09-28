package main

import (
	"fmt"
	"image"
)

func main() {
	fmt.Println(
		isPointInTriangle(
			image.Pt(1, 1),
			image.Pt(3, 1),
			image.Pt(1, 3),
			image.Pt(2, 2),
		),
	)
}

func isPointInTriangle(p1, p2, p3 image.Point, point image.Point) bool {
	return false
}
