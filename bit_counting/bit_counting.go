package bit_counting

import "math/bits"

//Write a function that takes an integer as input, and returns the number of bits that are equal to one in the binary representation of that number. You can guarantee that input is non-negative.
//
//Example: The binary representation of 1234 is 10011010010, so the function should return 5 in this case

func CountBits(n uint) int {
	return bits.OnesCount(n)
}

//func CountBits(n uint) int {
//	var res int = 0
//	for (n>0) {
//		if (n & 1 == 1) {
//			res = res + 1
//		}
//		n = n >> 1
//	}
//	return res
//}
