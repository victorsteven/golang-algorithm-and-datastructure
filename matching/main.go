package main

import (
	"fmt"
	"regexp"
)

//I got lots of files beginning like this:
//Program title: Primes
//Author: Kern
//Corporation: Gold
//Phone: +1-503-555-0091
//Date: Tues April 9, 2005
//Version: 6.7
//Level: Alpha

//Here we will work with strings like the string data above and not with files.
//The function change(s, prog, version) given:
//s=data, prog="Ladder" , version="1.1" will return:
//"Program: Ladder Author: g964 Phone: +1-503-555-0090 Date: 2019-01-01 Version: 1.1"

//Rules:
//The date should always be "2019-01-01".
//The author should always be "g964".
//
//Replace the current "Program Title" with the prog argument supplied to your function. Also remove "Title", so in the example case "Program Title: Primes" becomes "Program: Ladder".
//
//Remove the lines containing "Corporation" and "Level" completely.
//
//Phone numbers and versions must be in valid formats.
//
//A valid version in the given string data is one or more digits followed by a dot, followed by one or more digits. So 0.6, 5.4, 14.275 and 1.99 are all valid, but versions like .6, 5, 14.2.7 and 1.9.9 are invalid.
//
//A valid phone format is +1-xxx-xxx-xxxx, where each x is a digit.
//
//If the phone or version format is not valid, return "ERROR: VERSION or PHONE".
//
//If the version format is valid and the version is anything other than 2.0, replace it with the version parameter supplied to your function. If it’s 2.0, don’t modify it.
//
//If the phone number is valid, replace it with "+1-503-555-0090".
//
//Note
//You can see other examples in the "Sample tests".

func main() {

	s1 := "Program title: Primes\nAuthor: Kern\nCorporation: Gold\nPhone: +1-503-555-0091\nDate: Tues April 9, 2005\nVersion: 6.7\nLevel: Alpha"

	fmt.Println(Change(s1, "Ladder", "1.1"))

}

type rec struct {
	program string
	author  string
	phone   string
	date    string
	version string
}

func (r rec) String() string {
	return fmt.Sprintf("Program: %s Author: %s Phone: %s Date: %s Version: %s", r.program, r.author, r.phone, r.date, r.version)
}

func Change(s string, prog string, version string) string {
	r := rec{program: prog, author: "g964", date: "2019-01-01", phone: "+1-503-555-0090", version: "2.0"}
	re := regexp.MustCompile(`(.*)\: *(.*)\n?`)
	//fmt.Println("this is the string re: ", re)
	//fmt.Printf("the type of the re variable: %T\n", re)
	scan := re.FindAllStringSubmatch(s, -1)
	//fmt.Println("this is the scan result: ", scan)
	for _, l := range scan {
		//fmt.Println("this is the L value: ", l)
		//fmt.Println("this is the L0 value: ", l[0])
		//fmt.Println("this is the L1 value: ", l[1])
		//fmt.Println("this is the L2 value: ", l[2])

		switch l[1] {
		case "Phone":
			matched, _ := regexp.MatchString(`1-\d{3}-\d{3}-\d{4}`, l[2])
			if matched {
				return "ERROR: VERSION or PHONE"
			}
		case "Version":
			if l[2] == "2.0" {
				continue
			}
			matched, _ := regexp.MatchString(`\A\d{1,}\.{1}\d{1,}\z`, l[2])
			if !matched {
				return "ERROR: VERSION or PHONE"
			}
			r.version = version
		}
	}
	return r.String()
}

//Solution 2
func Change2(s string, prog string, version string) (ret string) {
	re, _ := regexp.Compile(`^Program title: .+\nAuthor: .+\nCorporation: .+\nPhone: \+1-\d{3}-\d{3}-\d{4}\nDate: .+\nVersion: (\d\. \d+)\nLevel: .+$`)
	ss := re.FindStringSubmatch(s)
	if len(ss) == 0 {
		return "ERROR: VERSION or PHONE"
	}
	v := "2.0"
	if ss[1] != "2.0" {
		v = version
	}
	ret += "Program: " + prog + " Author: g964 Phone: +1-503-555-0090 Date: 2019-01-01 Version: " + v
	return
}
