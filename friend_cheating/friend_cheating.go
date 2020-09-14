package friend_cheating


//A friend of mine takes a sequence of numbers from 1 to n (where n > 0).
//Within that sequence, he chooses two numbers, a and b.
//He says that the product of a and b should be equal to the sum of all numbers in the sequence, excluding a and b.
//Given a number n, could you tell me the numbers he excluded from the sequence?
//The function takes the parameter: n (n is always strictly greater than 0) and returns an array or a string (depending on the language) of the form:
//
//[(a, b), ...] or [[a, b], ...] or {{a, b}, ...} or or [{a, b}, ...]
//with all (a, b) which are the possible removed numbers in the sequence 1 to n.
//
//[(a, b), ...] or [[a, b], ...] or {{a, b}, ...} or ...will be sorted in increasing order of the "a".
//
//It happens that there are several possible (a, b). The function returns an empty array (or an empty string) if no possible numbers are found which will prove that my friend has not told the truth! (Go: in this case return nil).
//
//(See examples of returns for each language in "RUN SAMPLE TESTS")
//
//Examples:
//removNb(26) should return [(15, 21), (21, 15)]
//or
//
//removNb(26) should return { {15, 21}, {21, 15} }
//or
//
//removeNb(26) should return [[15, 21], [21, 15]]
//or
//
//removNb(26) should return [ {15, 21}, {21, 15} ]
//or
//
//removNb(26) should return "15 21, 21 15"
//or
//
//in C:
//removNb(26) should return  **an array of pairs {{15, 21}{21, 15}}**
//tested by way of strings.


func RemovNb(m uint64) [][2]uint64 {
	n := m
	var sum uint64 = n * (n + 1) / 2
	var res [][2]uint64
	for a := uint64(1); a <= n; a++ {
		if (sum - a) % (a + 1) == 0 {
			b := (sum - a) / (a + 1)
			if b < n {
				res = append(res, [2]uint64{a, b})
			}
		}
	}
	return res
}

