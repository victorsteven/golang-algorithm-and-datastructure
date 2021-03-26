package main

import "fmt"

//The Ulam sequence U is defined by u0=u, u1=v, with the general term u_n for n>2 given by the least integer expressible uniquely as the sum of two distinct earlier terms.
//In other words, the next number is always the smallest, unique sum of any two previous terms.
//
//The first 10 terms of the sequence U(u0=1, u1=2) are [1, 2, 3, 4, 6, 8, 11, 13, 16, 18].
//
//Here, the first term after the initial 1, 2 is obviously 3 since 3=1+2.
//The next term is 4=1+3 (we don't have to worry about 4=2+2 since it is a sum of a single term instead of two distinct terms)
//5 is not a member of the sequence since it is representable in two ways: 5=1+4=2+3, but 6=2+4 is a member (since 3+3 isn't a valid calculation).
//You'll have to write a code that creates an Ulam Sequence starting with u0, u1 and containing n terms.

// samples := [...][]int{
//        []int{1,2,5},
//        []int{3,4,5},
//        []int{5,6,8},
//        []int{3,4,5},
//        []int{1,2,20},
//        []int{1,3,60},
//    }
//    sample_sols := [...][]int{
//        []int{1,2,3,4,6},
//        []int{3,4,7,10,11},
//        []int{5,6,11,16,17,21,23,26},
//        []int{3,4,7,10,11},
//        []int{1,2,3,4,6,8,11,13,16,18,26,28,36,38,47,48,53,57,62,69},
//        []int{1,3,4,5,6,8,10,12,17,21,23,28,32,34,39,43,48,52,54,59,63,68,72,74,79,83,98,99,101,110,114,121,125,132,136,139,143,145,152,161,165,172,176,187,192,196,201,205,212,216,223,227,232,234,236,243,247,252,256,258},
//    }

func main() {
	result := UlamSequence(1, 2, 5)
	fmt.Println(result)
}

//func UlamSequence(u, v, n int) []int {
//
//	p1 := u
//	p2 := v
//
//	numbers := []int{p1, p2}
//
//	for i := v; i < n; i++ {
//		if p1 != p2 {
//
//		}
//		numbers = append(numbers, u + v)
//	}
//
//	return numbers
//}

func UlamSequence(n1, n2, q int) []int {

	tome := []int{n1, n2}

	codex := map[int]bool{n1 + n2: true}

	reject := make(map[int]bool)

	cv := n1 + n2

	c := 2

	for c < q {
		if codex[cv] {
			for _, n := range tome {
				cvn := cv + n
				if codex[cvn] {
					delete(codex, cvn)
					reject[cvn] = true
				} else if !reject[cvn] {
					codex[cvn] = true
				}
			}
			tome = append(tome, cv)
			c++
		}
		cv++
	}
	return tome
}
