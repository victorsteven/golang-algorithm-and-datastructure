package main

import (
	"fmt"
	"math"
)

//A rectangle with sides equal to even integers a and b is drawn on the Cartesian plane. Its center (the intersection point of its diagonals) coincides with the point (0, 0), but the sides of the rectangle are not parallel to the axes; instead, they are forming 45 degree angles with the axes.
//
//How many points with integer coordinates are located inside the given rectangle (including on its sides)?

func main() {

	ans := RectangleRotation2(6, 4)

	fmt.Println(ans)
}

func RectangleRotation2(a, b int) int {
	c := math.Sqrt(2) / 2
	x, y := int(float64(a)/2/c), int(float64(b)/2/c)
	return (x*2+1)*(y*2+1)/2 + (1 - ((x + y) % 2))
}
