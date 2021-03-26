package two_oldest_ages

import "sort"

//The two oldest ages function/method needs to be completed. It should take an array of numbers as its argument and return the two highest numbers within the array. The returned value should be an array in the format [second oldest age, oldest age].
//
//The order of the numbers passed in could be any order. The array will always include at least 2 items.
//
//For example:
//
//TwoOldestAges([]int{1, 5, 87, 45, 8, 8}) // should return [2]int{45, 87}

func TwoOldestAges(ages []int) [2]int {
	a, b := 0, 0
	for _, v := range ages {
		if v > b {
			a, b = b, v
		} else if v > a {
			a = v
		}
	}
	return [2]int{a, b}
}

func TwoOldestAges2(ages []int) [2]int {

	sort.Sort(sort.Reverse(sort.IntSlice(ages)))

	return [2]int{ages[1], ages[0]}
}

func TwoOldestAges3(ages []int) [2]int {

	sort.Ints(ages)

	first, second := ages[len(ages)-1], ages[len(ages)-2]

	return [2]int{second, first}

}
