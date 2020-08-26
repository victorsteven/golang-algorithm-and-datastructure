package example_3

import "fmt"

//In this example, we will group all alphabets and count all same alphabets

func Solution1(str string) string {

	mapStr := make(map[string]uint)

	result := ""

	for _, v := range str {
		mapStr[string(v)]++
	}

	for key, val := range mapStr {
		if val == 1 {
			result += fmt.Sprintf("%s", key)
		} else {
			result += fmt.Sprintf("%d%s", val, key)
		}
	}

	return result
}

