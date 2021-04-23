/*
Imagine we have an image. We'll represent this image as a simple 2D array where every pixel is a 1 or a 0. The image you get is known to have a single rectangle of 0s on a background of 1s.

Write a function that takes in the image and returns one of the following representations of the rectangle of 0's: top-left coordinate and bottom-right coordinate OR top-left coordinate, width, and height.

image1 = [
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 0, 0, 0, 1],
  [1, 1, 1, 0, 0, 0, 1],
  [1, 1, 1, 1, 1, 1, 1],
]

Sample output variations (only one is necessary):

findRectangle(image1) =>
  x: 3, y: 2, width: 3, height: 2
  2,3 3,5 -- row,column of the top-left and bottom-right corners

Other test cases:

image2 = [
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 0],
]

findRectangle(image2) =>
  x: 6, y: 4, width: 1, height: 1
  4,6 4,6

image3 = [
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 0, 0],
  [1, 1, 1, 1, 1, 0, 0],
]

findRectangle(image3) =>
  x: 5, y: 3, width: 2, height: 2
  3,5 4,6

image4 = [
  [0, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
  [1, 1, 1, 1, 1, 1, 1],
]

findRectangle(image4) =>
  x: 0, y: 0, width: 1, height: 1
  0,0 0,0

image5 = [
  [0],
]

findRectangle(image5) =>
  x: 0, y: 0, width: 1, height: 1
  0,0 0,0

n: number of rows in the input image
m: number of columns in the input image

*/

package main

import "fmt"

// TODO --- Write your function

func main() {

	//star(10)

	image1 := [][]int{
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 0, 0, 0, 1},
		[]int{1, 1, 1, 0, 0, 0, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
	}

	image2 := [][]int{
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 0},
	}

	image3 := [][]int{
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 0, 0},
		[]int{1, 1, 1, 1, 1, 0, 0},
	}

	image4 := [][]int{
		[]int{0, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
		[]int{1, 1, 1, 1, 1, 1, 1},
	}

	image5 := [][]int{
		[]int{0},
	}

	fmt.Println(rect(image1))
	fmt.Println(rect(image2))
	fmt.Println(rect(image3))
	fmt.Println(rect(image4))
	fmt.Println(rect(image5))

}

func rect(image [][]int) []int {

	theLengthOfRows := len(image)
	theLengthOfColumns := len(image[0])
	theSlice := []int{}

	for i := 0; i < theLengthOfRows; i++ {
		for j := 0; j < theLengthOfColumns; j++ {
			if image[i][j] == 1 {
				continue
			} else {
				pos := []int{}
				pos = append(pos, i)
				pos = append(pos, j)

				theSlice = append(theSlice, pos...)
			}
		}
	}

	x, y, w, z := theSlice[0], theSlice[1], theSlice[len(theSlice)-2], theSlice[len(theSlice)-1]

	theFinal := []int{}
	theFinal = append(theFinal, x)
	theFinal = append(theFinal, y)
	theFinal = append(theFinal, w)
	theFinal = append(theFinal, z)

	return theFinal

}
//
func star(n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func starReverse(n int) {

}
