package parts_of_a_list

import (
	"fmt"
	"regexp"
	"strings"
)

//Write a function partlist that gives all the ways to divide a list (an array) of at least two elements into two non-empty parts.
//
//Each two non empty parts will be in a pair (or an array for languages without tuples or a structin C - C: see Examples test Cases - )
//Each part will be in a string
//Elements of a pair must be in the same order as in the original array.
//Examples of returns in different languages:
//a = ["az", "toto", "picaro", "zone", "kiwi"] -->
//
//[["az", "toto picaro zone kiwi"], ["az toto", "picaro zone kiwi"], ["az toto picaro", "zone kiwi"], ["az toto picaro zone", "kiwi"]]
//
//or
//
//a = {"az", "toto", "picaro", "zone", "kiwi"} -->
//
//{{"az", "toto picaro zone kiwi"}, {"az toto", "picaro zone kiwi"}, {"az toto picaro", "zone kiwi"}, {"az toto picaro zone", "kiwi"}}
//
//or
//
//a = ["az", "toto", "picaro", "zone", "kiwi"] -->
//
//[("az", "toto picaro zone kiwi"), ("az toto", "picaro zone kiwi"), ("az toto picaro", "zone kiwi"), ("az toto picaro zone", "kiwi")]
//
//or
//
//a = [|"az", "toto", "picaro", "zone", "kiwi"|] -->
//
//[("az", "toto picaro zone kiwi"), ("az toto", "picaro zone kiwi"), ("az toto picaro", "zone kiwi"), ("az toto picaro zone", "kiwi")]
//
//or
//
//a = ["az", "toto", "picaro", "zone", "kiwi"] -->
//
//"(az, toto picaro zone kiwi)(az toto, picaro zone kiwi)(az toto picaro, zone kiwi)(az toto picaro zone, kiwi)"

func PartList(arr []string) string {

	str := ""

	for i := 1; i < len(arr); i++ {
		str += fmt.Sprintf("(%s, %s)", arr[:i], arr[i:])
	}

	reg := regexp.MustCompile(`[\[\]]`)
	strNew := reg.ReplaceAllString(str, "")

	return strNew
}

////////////////////////////////////////////////////////////////////////////
func join(arr []string) string {
	return strings.Join(arr, " ")
}

func PartList2(arr []string) string {

	var sb strings.Builder

	for i := 1; i < len(arr); i++ {
		fmt.Fprintf(&sb, "(%s, %s)", join(arr[:i]), join(arr[i:]))
	}

	return sb.String()
}
