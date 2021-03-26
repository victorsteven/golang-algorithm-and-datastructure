package ordered_count_characters

import (
	"strings"
)

//Count the number of occurrences of each character and return it as a list of tuples in order of appearance. For empty output return an empty list.

//Example:
//
//OrderedCount("abracadabra") == []Tuple{Tuple{'a', 5}, Tuple{'b', 2}, Tuple{'r', 2}, Tuple{'c', 1}, Tuple{'d', 1}}
//
//// Where
//type Tuple struct {
//    Char  rune
//    Count int
//}

type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(text string) (tups []Tuple) {

	str := ""

	for _, v := range text {
		if strings.Count(str, string(v)) == 0 {
			tups = append(tups, Tuple{
				Char:  v,
				Count: strings.Count(text, string(v)),
			})
			str += string(v)
		}
	}
	return tups
}

//A different approach
func OrderedCount2(text string) (tups []Tuple) {

	var hashMap = make(map[rune]int)
	for _, v := range text {
		if _, ok := hashMap[v]; !ok {
			count := strings.Count(text, string(v))
			hashMap[v] = count

			tups = append(tups, Tuple{
				Char:  v,
				Count: count,
			})
		}
	}
	return tups
}
