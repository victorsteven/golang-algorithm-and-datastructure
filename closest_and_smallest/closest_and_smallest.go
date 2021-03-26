package closest_and_smallest

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

//a string string of n positive numbers (n = 0 or n >= 2)
//Let us call weight of a number the sum of its digits. For example 99 will have "weight" 18, 100 will have "weight" 1.
//
//Two numbers are "close" if the difference of their weights is small.
//
//Task:
//For each number in strng calculate its "weight" and then find two numbers of strng that have:
//
//the smallest difference of weights ie that are the closest
//with the smallest weights
//and with the smallest indices (or ranks, numbered from 0) in strng
//Output:
//
//an array of two arrays, each subarray in the following format:
//[number-weight, index in strng of the corresponding number, original corresponding number instrng]
//
//or a pair of two subarrays (Haskell, Clojure, FSharp) or an array of tuples (Elixir, C++)
//
//or a (char*) in C mimicking an array of two subarrays or a string
//
//or a matrix in R (2 rows, 3 columns, no columns names)
//
//The two subarrays are sorted in ascending order by their number weights if these weights are different, by their indexes in the string if they have the same weights.
//
//Examples:
//Let us call that function closest
//
//strng = "103 123 4444 99 2000"
//the weights are 4, 6, 16, 18, 2 (ie 2, 4, 6, 16, 18)
//
//closest should return [[2, 4, 2000], [4, 0, 103]] (or ([2, 4, 2000], [4, 0, 103])
//or [{2, 4, 2000}, {4, 0, 103}] or ... depending on the language)
//because 2000 and 103 have for weight 2 and 4, their indexes in strng are 4 and 0.
//The smallest difference is 2.
//4 (for 103) and 6 (for 123) have a difference of 2 too but they are not
//the smallest ones with a difference of 2 between their weights.
//....................
//
//strng = "80 71 62 53"
//All the weights are 8.
//closest should return [[8, 0, 80], [8, 1, 71]]
//71 and 62 have also:
//- the smallest weights (which is 8 for all)
//- the smallest difference of weights (which is 0 for all pairs)
//- but not the smallest indices in strng.
//....................
//
//strng = "444 2000 445 544"
//the weights are 12, 2, 13, 13 (ie 2, 12, 13, 13)
//
//closest should return [[13, 2, 445], [13, 3, 544]] or ([13, 2, 445], [13, 3, 544])
//or [{13, 2, 445}, {13, 3, 544}] or ...
//444 and 2000 have the smallest weights (12 and 2) but not the smallest difference of weights;
//they are not the closest.
//Here the smallest difference is 0 and in the result the indexes are in ascending order.
//...................
//
//closest("444 2000 445 644 2001 1002") --> [[3, 4, 2001], [3, 5, 1002]] or ([3, 4, 2001],
//[3, 5, 1002]]) or [{3, 4, 2001}, {3, 5, 1002}] or ...
//Here the smallest difference is 0 and in the result the indexes are in ascending order.
//...................
//
//closest("239382 162 254765 182 485944 468751 49780 108 54")
//The weights are: 27, 9, 29, 11, 34, 31, 28, 9, 9.
//closest should return  [[9, 1, 162], [9, 7, 108]] or ([9, 1, 162], [9, 7, 108])
//or [{9, 1, 162}, {9, 7, 108}] or ...
//108 and 54 have the smallest difference of weights too, they also have
//the smallest weights but they don't have the smallest ranks in the original string.
//..................
//
//closest("54 239382 162 254765 182 485944 468751 49780 108")
//closest should return  [[9, 0, 54], [9, 2, 162]] or ([9, 0, 54], [9, 2, 162])
//or [{9, 0, 54}, {9, 2, 162}] or ...
//Notes :
//If n == 0 closest("") should return []
//
//or ([], []) in Haskell, Clojure, FSharp
//
//or [{}, {}] in Elixir or '(() ()) in Racket
//
//or {{0,0,0}, {0,0,0}} in C++
//
//or "[(), ()]" in Go, Nim,
//
//or "{{0,0,0}, {0,0,0}}" in C, NULL in R.
//
//See Example tests for the format of the results in your language.

type Number struct {
	DigiStr  string
	Weight   int
	Position int
}

func Closest(str string) string {

	if len(str) == 0 {
		return "[(), ()]"
	}

	strArr := strings.Split(str, " ")

	numbers := make([]Number, len(strArr))

	for i, v := range strArr {
		weight := 0
		for _, x := range v {
			vint, _ := strconv.Atoi(string(x))
			weight += vint
		}

		var number = Number{
			DigiStr:  v,
			Weight:   weight,
			Position: i,
		}

		numbers[i] = number

	}

	//Sort the number by weight and position
	sort.Slice(numbers, func(i, j int) bool {
		var number1 = numbers[i]
		var number2 = numbers[j]

		//if the weight is different, sort by weight:
		if number1.Weight != number2.Weight {
			return number1.Weight < number2.Weight
		}

		//if the weight is the same, sort by position
		return number1.Position < number2.Position
	})

	//loop through each element and find the position of 2 items with smallest different of weight
	var minPos1, minPos2, minDiff = -1, -1, math.MaxInt64

	for i := 1; i < len(numbers); i++ {
		var diff = numbers[i].Weight - numbers[i-1].Weight
		if diff < minDiff {
			minDiff = diff
			minPos1 = i - 1
			minPos2 = i
		}
	}

	//build the result
	var number1 = numbers[minPos1]
	var number2 = numbers[minPos2]

	return fmt.Sprintf("[(%d, %d, %s), (%d, %d, %s)]", number1.Weight, number1.Position, number1.DigiStr, number2.Weight, number2.Position, number2.DigiStr)

}

///////////////////////////////////////////////////////////////////////////////////////
type Num struct {
	weight, index, num int
}

type NumberList []Num

func (a NumberList) Len() int {
	return len(a)
}

func (a NumberList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a NumberList) Less(i, j int) bool {

	//if the weight are not equal, sort by the weight
	if a[i].weight != a[j].weight {
		return a[i].weight < a[j].weight
	}

	//else, sort by the positon
	return a[i].index < a[j].index
}

func DigitSum(x int) int {
	res := 0
	for ; x > 0; x /= 10 {
		res += x % 10
	}
	return res
}

func NumberToString(a Num) string {
	return fmt.Sprintf("(%d, %d, %d)", a.weight, a.index, a.num)
}

func Closest2(str string) string {

	var numbers NumberList

	for idx, s := range strings.Split(str, " ") {
		var number Num
		number.num, _ = strconv.Atoi(s)
		number.index = idx
		number.weight = DigitSum(number.num)
		numbers = append(numbers, number)
	}

	sort.Sort(numbers)

	var num1, num2 Num

	min := math.MaxInt32

	for i, _ := range numbers {
		if i == 0 {
			continue
		}

		tmp := numbers[i].weight - numbers[i-1].weight

		if tmp < min {
			min = tmp
			num1 = numbers[i-1]
			num2 = numbers[i]
		}
	}

	if min == math.MaxInt32 {
		return "[(), ()]"
	}

	return fmt.Sprintf("[%s, %s]", NumberToString(num1), NumberToString(num2))
}
