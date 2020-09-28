package main

import (
	"fmt"
	"image"
	"math"

	"github.com/fr3fou/matrigo"
)

func main() {
	fmt.Println(
		isPointInTriangleDet(
			image.Pt(1, 1),
			image.Pt(3, 1),
			image.Pt(1, 3),
			image.Pt(2, 2),
		),
	)
	fmt.Println(
		isPointInTriangleDet(
			image.Pt(1, 1),
			image.Pt(3, 1),
			image.Pt(1, 3),
			image.Pt(2, 4),
		),
	)
}

func isPointInTriangleDet(p1, p2, p3 image.Point, p image.Point) bool {
	mainArea := matrigo.New(3, 3, [][]float64{
		{float64(p1.X), float64(p1.Y), 1},
		{float64(p2.X), float64(p2.Y), 1},
		{float64(p3.X), float64(p3.Y), 1},
	}).Det() / 2

	sum := 0.0

	sum += math.Abs(matrigo.New(3, 3, [][]float64{
		{float64(p.X), float64(p.Y), 1},
		{float64(p2.X), float64(p2.Y), 1},
		{float64(p3.X), float64(p3.Y), 1},
	}).Det() / 2)

	sum += math.Abs(matrigo.New(3, 3, [][]float64{
		{float64(p1.X), float64(p1.Y), 1},
		{float64(p.X), float64(p.Y), 1},
		{float64(p3.X), float64(p3.Y), 1},
	}).Det() / 2)

	sum += math.Abs(matrigo.New(3, 3, [][]float64{
		{float64(p1.X), float64(p1.Y), 1},
		{float64(p2.X), float64(p2.Y), 1},
		{float64(p.X), float64(p.Y), 1},
	}).Det() / 2)

	return mainArea == sum
}
