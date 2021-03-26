package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"A", "B", "C", "D"}
	fmt.Println(StockList(b, c))
}

func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}
	result := ""
	for _, v := range listCat {
		tot := 0
		for _, w := range listArt {
			//check the first character for the listCat and the listArt slices,
			//w[0] fetches the first character ASCII value, same with v[0],
			//so if listCat = []string{"A", "B", "C", "D"}, then, v[0] = 65, which corresponds to A
			//if we have listArt = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}, then w[0]= 66, which corresponds to B
			//so the comparison goes like this: if 65 == 66 (depending on the iteration)
			if w[0] == v[0] {
				ans, _ := strconv.Atoi(strings.Split(w, " ")[1])
				tot += ans
			}
		}
		if result != "" {
			result += " - "
		}
		result += "(" + v + " : " + strconv.Itoa(tot) + ")"
	}
	return result
}
