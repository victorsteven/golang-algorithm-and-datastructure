package twice_linear

//Consider a sequence u where u is defined as follows:
//
//The number u(0) = 1 is the first one in u.
//For each x in u, then y = 2 * x + 1 and z = 3 * x + 1 must be in u too.
//There are no other numbers in u.
//Ex: u = [1, 3, 4, 7, 9, 10, 13, 15, 19, 21, 22, 27, ...]
//
//1 gives 3 and 4, then 3 gives 7 and 10, 4 gives 9 and 13, then 7 gives 15 and 22 and so on...
//
//Task:
//Given parameter n the function dbl_linear (or dblLinear...) returns the element u(n) of the ordered (with <) sequence u (so, there are no duplicates).
//
//Example:
//dbl_linear(10) should return 22
//
//Note:
//Focus attention on efficiency

func DblLinear(n int) int {
	// your code
	u := []int{1}
	i := 0
	j := 0

	var y int
	var z int

	for len(u) <= n {
		y = 2*u[i] + 1
		z = 3*u[j] + 1

		if y < z {
			u = append(u, y)
			i++
		} else if y > z {
			u = append(u, z)
			j++
		} else {
			u = append(u, y)
			i++
			j++
		}
	}
	//fmt.Println(u)
	return u[len(u)-1]
}

func DblLinear2(n int) int {

	u := make([]int, 0)

	u = append(u, 1)
	x, y := 0, 0

	for i := 0; i < n; i++ {

		next := 2*u[x] + 1
		neyt := 3*u[y] + 1

		if next <= neyt {
			u = append(u, next)
			x++
			if next == neyt {
				y++
			}
		} else {
			u = append(u, neyt)
			y++
		}
	}
	return u[n]
}
