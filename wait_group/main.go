package main

import (
	"fmt"
	"strings"
)

func main() {
	boysData := "Mike 3rd, Victor 2nd, Kalu 1st, Ode 4th, Grand 5th"
	fmt.Println(boys(boysData))
}

func boys(b string) map[string]string {
	ans := make(map[string]string)
	bSlice := strings.Split(b, ",")
	for _, v := range bSlice {
		w := strings.Split(v, " ")
		ans[w[0]] = w[1]
	}
	return ans
}
