package main

import "fmt"

func main() {
	fmt.Println("this is the ans: ", CreateSpiral(5))
}

//func CreateSpiral(n int) [][]int {
//	if n < 1 {
//		return [][]int{}
//	}
//	var matirx [][]int
//
//	top, left, bottom, right := 0,0, n-1, n-1
//	sz := n * n
//	a := make([]int, sz)
//	i := 0
//	for left < right {
//		//	work right, along top
//		for c := left; c <= right; c++ {
//			a[top*n+c] = i
//			i++
//		}
//		top++
//		//work down right side
//		for r := top; r <= bottom; r++ {
//			a[r*n+right] = i
//			i++
//		}
//		right--
//		if top == bottom {
//			break
//		}
//		//Work left, along bottom
//		for c := right; c >= left; c-- {
//			a[bottom*n+c] = i
//			i++
//		}
//		bottom--
//
//		//work up left side
//		for r := bottom; r >= top; r-- {
//			a[r*n+left] = i
//			i++
//		}
//		left++
//	}
//	a[top*n+left] = i
//
//	for m := 0; m <= n - 1; m++ {
//		l := m*n
//		h := m*n+n
//		matirx = append(matirx, a[l:h])
//	}
//	return matirx
//}

func CreateSpiral(n int) [][]int {
	// your code here
	var matrice [][]int
	if n > 0 {
		l := 0
		t := 0
		r := n-1
		b := n-1

		line := make([]int, n * n)
		i := 1
		for l < r {
			for c := l; c <= r; c++ {
				line[t*n+c] = i
				i++
			}
			t++
			for z := t; z <= b; z++ {
				line[z*n+r] = i
				i++
			}
			r--
			if t == b {
				break
			}
			for c := r; c >= l; c-- {
				line[b*n+c] = i
				i++
			}
			b--
			for z := b; z >= t; z-- {
				line[z*n+l] = i
				i++
			}
			l++
		}
		line[t*n+l] = i
		// return matrice
		for m := 0; m <= n-1; m++ {
			low := m*n
			high := m*n+n
			matrice = append(matrice,line[low:high])
		}
		return matrice
	}
	return [][]int{}

}