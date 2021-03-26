package main

import (
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(rangeFormat([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))

	fmt.Println(Solu([]int{-6, -3, -2, -1, 0, 1, 3, 4, 5, 7, 8, 9, 10, 11, 14, 15, 17, 18, 19, 20}))
}

func Solution(list []int) string {
	//for _, v := range list {
	//	fmt.Println(v)
	//}
	var str string
	for i := 0; i < len(list); i++ {
		if list[i]+1 != list[i+1] {
			str += strconv.Itoa(list[i]) + ","
		} else if (list[i]+1) == list[i+1] && (list[i+1]+1) == list[i+2] && list[i-1] != list[i]-1 {
			str += strconv.Itoa(list[i]) + "-"
		} else if strconv.Itoa(int(str[len(str)-1])) == "-" && list[i+1] != list[i]+1 {
			str += strconv.Itoa(list[i]) + ","
		} else if list[i]-1 == list[i-1] && list[i]+1 == list[i+1] {

		} else if list[i]-1 == list[i-1] && list[i-1]-1 == list[i-2] && list[i-1] != list[i]-1 {
			str += strconv.Itoa(list[i]) + ","
		} else {
			str += strconv.Itoa(list[i]) + ","
		}
	}
	return str[0:]
}

func rangeFormat(list []int) (string, error) {
	if len(list) == 0 {
		return "", nil
	}
	var parts []string
	for n1 := 0; ; {
		n2 := n1 + 1
		for n2 < len(list) && list[n2] == list[n2-1]+1 {
			n2++
		}
		s := strconv.Itoa(list[n1])
		if n2 == n1+2 {
			s += "," + strconv.Itoa(list[n2-1])
		} else if n2 > n1+2 {
			s += "-" + strconv.Itoa(list[n2-1])
		}
		parts = append(parts, s)
		if n2 == len(list) {
			break
		}
		if list[n2] == list[n2-1] {
			return "", errors.New(fmt.Sprintf("sequence repeats value %d", list[n2]))
		}
		if list[n2] < list[n2-1] {
			return "", errors.New(fmt.Sprintf("sequence is not ordered: %d < %d", list[n2], list[n2-1]))
		}
		n1 = n2
	}
	return strings.Join(parts, ","), nil
}

func Solu(list []int) (s string) {
	l := len(list) - 1
	for i, j := 0, 0; i < l; i = j {
		s += strconv.Itoa(list[i])
		for j = i; (j < l) && (list[j]+1 == list[j+1]); {
			j++
		}
		if j-i > 1 {
			s += "-"
		} else {
			s += ","
		}
		if i == j {
			j++
		}
	}
	s += strconv.Itoa(list[l])
	return
}

func Solu2(list []int) string {
	var res bytes.Buffer
	for i := 0; i < len(list); i++ {
		if i != 0 {
			res.WriteRune(',')
		}
		j := i
		for j+1 < len(list) && list[j]+1 == list[j+1] {
			j++
		}
		if i+1 == j {
			j = i
		}
		res.WriteString(fmt.Sprint(list[i]))
		if i != j {
			res.WriteRune('-')
			res.WriteString(fmt.Sprint(list[j]))
		}
		i = j
	}
	return res.String()
}
