package string_to_int

import "strings"

func atoi(s string) (int, error) {
	var n []int
	origString := s
	if s[0] == '-' {
		s = s[1:]
	}
	for _, v := range s {
		switch v {
		case '1','2','3','4','5','6','7','8','9','0':
			n = append(n, int(v - '0'))
		default:
			return 0, nil
		}
	}
	res := 0
	op := 1
	for i := len(n) - 1; i >=0; i-- {
		res += n[i] * op
		op *= 10
	}

	if strings.Contains(origString, "-") {
		res = res * -1
	}
	return res, nil
}
