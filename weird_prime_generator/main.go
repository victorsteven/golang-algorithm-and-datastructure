package main


//Consider the sequence a(1) = 7, a(n) = a(n-1) + gcd(n, a(n-1)) for n >= 2:
//
//7, 8, 9, 10, 15, 18, 19, 20, 21, 22, 33, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 69, 72, 73....
//
//Let us take the differences between successive elements of the sequence and get a second sequence g: 1, 1, 1, 5, 3, 1, 1, 1, 1, 11, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 23, 3, 1....
//
//For the sake of uniformity of the lengths of sequences we add a 1 at the head of g:
//
//g: 1, 1, 1, 1, 5, 3, 1, 1, 1, 1, 11, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 23, 3, 1...
//
//Removing the 1s gives a third sequence: p: 5, 3, 11, 3, 23, 3... where you can see prime numbers.
//
//#Task: Write functions:
//
//1: an(n) with parameter n: returns the first n terms of the series a(n) (not tested)
//
//2: gn(n) with parameter n: returns the first n terms of the series g(n) (not tested)
//
//3: countOnes(n) with parameter n: returns the number of 1 in g(n)
//(don't forget to add a `1` at the head) # (tested)
//
//4: p(n) with parameter n: returns an array of n unique prime numbers (not tested)
//
//5: maxp(n) with parameter n: returns the biggest prime number of the sequence pn(n) # (tested)
//
//6: anOver(n) with parameter n: returns an array (n terms) of the a(i)/i for every i such g(i) != 1 (not tested but interesting result)
//
//7: anOverAverage(n) with parameter n: returns as an *integer* the average of anOver(n)  (tested)



var (
	aCache    = map[int64]int64{1:7}
	pCache    []int64
	nextIndex = int64(1)
)

func hasPi(pi int64) bool {
	for _, v := range pCache {
		if v == pi { return true }
	}

	return false
}

func gcd(a, b int64) int64 {
	if a < b { a, b = b, a }
	for ; (a % b) != 0; a, b = b, a % b {}

	return b
}

func a(n int64) int64 {
	res, ok := aCache[n]
	if !ok {
		prev := a(n-1)
		res = prev + gcd(n, prev)

		aCache[n] = res
	}

	return res
}

func g(n int64) int64 {
	if n == 1 { return 1 }

	return a(n) - a(n-1)
}

func p(n int64) int64 {
	n--
	for ; n >= int64(len(pCache)); nextIndex++ {
		for ; g(nextIndex) == 1; nextIndex++ {}
		if !hasPi(g(nextIndex)) { pCache = append(pCache, g(nextIndex)) }
	}

	return pCache[n]
}

func CountOnes(n int64) (res int) {
	for i := int64(1); i <= n; i++ { if g(i) == 1 { res++ } }

	return
}

func MaxPn(n int64) (res int64) {
	for i := int64(1); i <= n; i++ { if res < p(i) { res = p(i) } }

	return
}

func AnOverAverage(n int64) int {
	sum := float64(0)
	idx := int64(1)
	for i := int64(1); i <= n; i++ {
		for ; g(idx) == 1; idx++ {}
		sum += float64(a(idx)) / float64(idx)
	}

	return int(sum / float64(n))
}




//SOLUTION 2
//func gcd(a, b int64) int64 {
//	if a < b { a, b = b, a }
//	for b > 0 { a, b = b, a % b }
//	return a
//}
//
//func CountOnes(n int64) (c int) {
//	for i, last := int64(2), int64(7); i <= n; i++ {
//		item := last + gcd(i, last)
//		if item - last == 1 { c++ }
//		last = item
//	}
//	c++
//	return
//}
//func MaxPn(n int64) (max int64) {
//	m := map[int64]int{}
//	for i, j, last := int64(2), int64(0), int64(7); j < n; i++ {
//		item := last + gcd(i, last)
//		k := item - last
//		if k != 1 && m[k] == 0 { m[k], j = 1, j + 1 }
//		if k > max { max = k }
//		last = item
//	}
//	return
//}
//func AnOverAverage(n int64) int {
//	sum := int64(0)
//	for i, j, last := int64(2), int64(0), int64(7); j < n; i++ {
//		item := last + gcd(i, last)
//		if item - last == 1 {
//			last = item
//			continue
//		}
//		sum, last, j = sum + item / i, item, j + 1
//	}
//	return int(sum / n)
//}




//SOLUIION
//func gcd(a, b int64) int64 {
//	if a < b {
//		a, b = b, a
//	}
//	for b > 0 {
//		a, b = b, a%b
//	}
//	return a
//}
//
//func CountOnes(n int64) int {
//	an, gn := int64(7), int64(1)
//	cnt := 1
//	for i := int64(2); i <= n; i++ {
//		gn = gcd(i, an)
//		an += gn
//		if gn == 1 {
//			cnt++
//		}
//	}
//	return cnt
//}
//
//func MaxPn(n int64) int64 {
//	an, gn, max := int64(7), int64(1),int64(0)
//	p := make(map[int64]bool)
//
//	for i := int64(2); int64(len(p)) < n; i++ {
//		gn = gcd(i, an)
//		an += gn
//		if gn > 1 {
//			p[gn] = true
//			if gn > max {
//				max = gn
//			}
//		}
//	}
//	return max
//}
//
//func AnOverAverage(n int64) int {
//	return 3  // Because a(i)/i = 3, always
//}