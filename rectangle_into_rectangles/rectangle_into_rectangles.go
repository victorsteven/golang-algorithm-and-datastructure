package rectangle_into_rectangles

import "fmt"

//A "True Rectangle" is a rectangle with two different dimensions and four equal angles.
//
//Task:
//In this kata, we want to decompose a given true rectangle into the minimum number of squares, Then aggregate these generated squares together to form all the possible true rectangles.

//Example: Refer to figure_1
//As shown in this figure, we want to decompose the (13*5) true rectangle into the minimum number of squares which are [5, 5, 3, 2, 1, 1] to return all the possible true rectangles from aggregating these squares together :

//rectIntoRects(13, 5) should return the ractangles: ["(10*5)", "(8*5)", "(2*1)", "(3*2)", "(5*3)", "(13*5)"]   //or any other order

//Example 2: Refer to figure_2
//Here is the (22*6) true rectangle, it will be decomposed into the [6, 6, 6, 4, 2, 2] squares. so we should aggregate these squares together to form all the possible true rectangles :
//rectIntoRects(22, 6) should return the ractangles: ["(12*6)", "(18*6)", "(22*6)", "(12*6)", "(16*6)", "(10*6)", "(6*4)", "(4*2)"]   //or any other order

//More Examples :
//
//The (8*5) true rectangle will be decomposed into [5, 3, 2, 1, 1] squares, so :
//
//rectIntoRects(8, 5) should return: ["(8*5)", "(5*3)", "(3*2)", "(2*1)"]   //or any other order
//The (20*8) rectangle will be decomposed into [8, 8, 4, 4] squares, so :
//
//rectIntoRects(20, 8) should return: ["(16*8)", "(20*8)", "(12*8)", "(8*4)"]   //or any other order
//See the example test cases for more examples.
//
//Notes:
//You should take each square with its all adjacent squares or rectangles to form the resulting true rectangles list.
//Do not take care of the resulting rectangles' orientation. just "(long_side*short_side)".
//Edge cases:
//rectIntoRects(17, 5) should equal rectIntoRects(5, 17).
//If length == width it should return null/None/{} based on your language.
//If length == 0 or width == 0 it should return null/None/{} based on your language.
//References:
//https://www.codewars.com/kata/55466989aeecab5aac00003e

func RectIntoRectangle1(l, w int) []string {

	l0, w0 := l, w

	lw := func() {
		if l < w {
			t := l
			l = w
			w = t
		}
	}
	sqs, rects := map[int]int{}, []string{}

	for l > 0 && w > 0 {
		lw()
		sqs[w]++
		l -= w
	}
	for sq, r := range sqs {
		for j := 2; j <= r; j++ {
			for k := 1; k <= 1+r-j; k++ {
				rects = append(rects, fmt.Sprintf("(%d*%d)", sq*j, sq))
			}
		}
	}

	l, w = l0, w0

	for l%w > 0 && w%l > 0 {
		lw()
		rects = append(rects, fmt.Sprintf("(%d*%d)", l, w))
		l -= w
	}
	return rects
}

////////////////////////////////////////////////////////////////////////////////////

func RectIntoRectangle2(a, b int) []string {
	s := squares(a, b)
	return rectangles(s)
}

func squares(long, short int) (ans []int) {
	for {
		if short > long {
			long, short = short, long
		}
		if short == 0 {
			return ans
		}
		ans = append(ans, short)

		long -= short
	}
}

func rectangles(squares []int) []string {

	ans := []string{}
	for i, short := range squares {
		j := i
		for j < len(squares) && squares[j] == squares[i] {
			j++
		}

		if j+1 < len(squares) {
			j++
		}

		long := short
		for k := i + 1; k < j; k++ {
			long += squares[k]
			s := fmt.Sprintf("(%d*%d)", long, short)
			ans = append(ans, s)
		}
	}

	return ans

}

func RectIntoRectangle3(l, w int) []string {

	res := []string{}

	if l == w || l == 0 || w == 0 {
		return res
	}

	for l > 0 {
		if w > l {
			w, l = l, w
		}
		if l != w {
			res = append(res, fmt.Sprintf("(%d*%d)", l, w))
		}
		i := l / w
		if l%w == 0 {
			i -= 1
		}

		for ; i > 1; i-- {
			res = append(res, fmt.Sprintf("(%d*%d)", w*i, w))
		}
		l -= w
	}

	return res

}

func RectIntoRectangle4(l, w int) []string {
	if l == w || l == 0 || w == 0 {
		return []string{}
	}
	res := make([]string, 0)
	for {
		if l == w {
			break
		}
		var long, short int
		if l > w {
			long, short = l, w
			l -= w
		} else {
			long, short = w, l
			w -= l
		}
		side := short * 2
		for side < long {
			res = append(res, fmt.Sprintf("(%d*%d)", side, short))
			side += short
		}
		if side >= long {
			res = append(res, fmt.Sprintf("(%d*%d)", long, short))
		}
	}
	return res
}
