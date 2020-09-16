package variadic_sum

//Sum function sums any number of elements given to it:

func Sum(nums ...int) int {

	total := 0

	for _, num := range nums {

		total += num
	}
	return total
}
