package geometry

import "math"

type Point struct {
	x float64
	y float64
}

func differenceBetweenPoints(a Point, b Point) float64 {
	var difX = math.Abs(a.x - b.x)
	var difY = math.Abs(a.y - b.y)
	return math.Sqrt(math.Pow(difX, 2) + math.Pow(difY, 2))
}
