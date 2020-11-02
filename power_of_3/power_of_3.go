package power_of_3

import (
	"math"
)

//Given a positive integer N, return the largest integer k such that 3^k < N.
//
//For example,
//
//LargestPower(3) // returns 0
//LargestPower(4) // returns 1

//You may assume that the input to your function is always a positive integer.

func LargestPower(n uint64) int {

	largest := 0

	for i := uint64(0); i < n; i++ {
		if math.Pow(3, float64(i)) < float64(n) && i > uint64(largest) {
			largest = int(i)
		}
	}
	return largest
}

//more performant solution:
func LargestPower2(n uint64) int {

	k := 0.0

	for math.Pow(3, k) < float64(n) {
		k++
	}

	return int(k - 1)
}


//import . "math"
//
//func LargestPower(n uint64) int {
//   k := 0.0
//   for Pow(3,k) < float64(n) {
//       k++
//   }
//   return int(k-1)
//}



//import "math"
//
//func LargestPower(n uint64) int {
//  k := 0
//
//  for uint64(math.Pow(3, float64(k))) < n {
//    k++
//  }
//
//  return k - 1
//}