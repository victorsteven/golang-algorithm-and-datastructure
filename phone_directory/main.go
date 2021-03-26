package main

import (
	"fmt"
	"regexp"
	"strings"
)

var dr = "/+1-541-754-3010 156 Alphand_St. <J Steeve>\n 133, Green, Rd. <E Kustur> NY-56423 ;+1-541-914-3010\n" + "+1-541-984-3012 <P Reed> /PO Box 530; Pollocksville, NC-28573\n :+1-321-512-2222 <Paul Dive> Sequoia Alley PQ-67209\n" + "+1-741-984-3090 <Peter Reedgrave> _Chicago\n :+1-921-333-2222 <Anna Stevens> Haramburu_Street AA-67209\n" + "+1-111-544-8973 <Peter Pan> LA\n +1-921-512-2222 <Wilfrid Stevens> Wild Street AA-67209\n" + "<Peter Gone> LA ?+1-121-544-8974 \n <R Steell> Quora Street AB-47209 +1-481-512-2222\n" + "<Arthur Clarke> San Antonio $+1-121-504-8974 TT-45120\n <Ray Chandler> Teliman Pk. !+1-681-512-2222! AB-47209,\n" + "<Sophia Loren> +1-421-674-8974 Bern TP-46017\n <Peter O'Brien> High Street +1-908-512-2222; CC-47209\n" + "<Anastasia> +48-421-674-8974 Via Quirinal    Roma\n <P Salinger> Main Street, +1-098-512-2222, Denver\n" + "<C Powel> *+19-421-674-8974 Chateau des Fosses Strasbourg F-68000\n <Bernard Deltheil> +1-498-512-2222; Mount Av.  Eldorado\n" + "+1-099-500-8000 <Peter Crush> Labrador Bd.\n +1-931-512-4855 <William Saurin> Bison Street CQ-23071\n" + "<P Salinge> Main Street, +1-098-512-2222, Denve\n" + "<P Salinge> Main Street, +1-098-512-2222, Denve\n"

func main() {
	//fmt.Println(Phone(dr, ""))
	//
	//fmt.Println(Phone2(dr, ""))

	//Phone3(s, "1-541-754-3010")
	s := "/+1-541-754-3010 156 Alphand_St. <J Steeve>\n 133, Green, Rd. <E Kustur> NY-56423 ;+1-541-914-3010!\n"

	//fmt.Println(Phone4(s, "1-541-754-3010"))
	//fmt.Println(Phone5(s, "1-541-754-3010"))
	fmt.Println(phone6(s, "1-541-754-3010"))

	//Phone3(dr, )

	//var regName = regexp.MustCompile(`<(.+)>\s`)
	//var regEnd = regexp.MustCompile(`[A-Za-z]+$`)
	//var regBegin = regexp.MustCompile(`^[A-Za-z]+`)
	//var regNegate = regexp.MustCompile(`[^A-Za-z]+`)

	//ans := regName.FindStringSubmatch("normal_http_call <frank> how are you?")
	//fmt.Println(regEnd.FindStringSubmatch("normal_http_call <frank> how are you"))
	//fmt.Println(regBegin.FindStringSubmatch("normal_http_call <frank> how are you"))
	//fmt.Println(regNegate.FindStringSubmatch("normal_http_call <!@#frank> how are you"))
	//fmt.Println(ans[1])
	//fmt.Println(ans2)

	//matched, _ := regexp.MatchString(`a.b`, "aaxbb")
	//fmt.Println(matched) // true
}

func Phone(dir, num string) string {
	cnt := 0
	data := ""
	for _, dd := range strings.Split(dir, "\n") {
		if strings.Contains(dd, "+"+num) {
			cnt++
			data = dd
		}
	}
	if cnt == 0 {
		return "Error => Not found: " + num
	} else {
		if cnt > 1 {
			return "Error => Too many people: " + num
		}
	}
	reg, _ := regexp.Compile("<|>")
	name := reg.Split(data, -1)[1]
	reg, _ = regexp.Compile(num + "|<.*>|[^A-Za-z0-9\\. -]")
	address := string(reg.ReplaceAll([]byte(data), []byte(" ")))
	reg, _ = regexp.Compile("\\s+")
	address = string(reg.ReplaceAll([]byte(address), []byte(" ")))
	address = strings.Trim(address, " ")
	return "Phone => " + num + ", Name => " + name + ", Address => " + address
}

func Phone2(dir, num string) string {
	c := strings.Count(dir, "+"+num)
	if c > 1 {
		return "Error => Too many people: " + num
	}
	if c == 0 {
		return "Error => Not found: " + num
	}
	ss := strings.Split(dir, "\n")
	ret := "Phone => " + num
	for _, v := range ss {
		i := strings.Index(v, "+"+num)
		for i < 0 {
			continue
		}
		v = v[:i] + v[i+len("+"+num):]
		i, j := strings.Index(v, "<"), strings.Index(v, ">")
		ret += ", Name => " + v[i+1:j]
		v = v[:i] + v[j+1:]
		for i, v1 := range v {
			if !(v1 >= 'A' && v1 <= 'Z' || v1 >= 'a' && v1 <= 'z' || v1 >= '0' && v1 <= '9' || v1 == '-' || v1 == '.') {
				v = v[:1] + " " + v[i+1:]
			}
		}
		return ret + ", Address => " + strings.Join(strings.Fields(v), " ")
	}
	return ""
}

func Phone3(dir, num string) string {
	var regPhone = regexp.MustCompile(`\+(\d{1,2}-\d{3}-\d{3}-\d{4})`)
	var regName = regexp.MustCompile(`<([^>]+)>\s?`) //can also be: `<(.+)>`
	var regSym = regexp.MustCompile(`[^\-.0-9A-Za-z]+`)

	ret := ""
	for _, line := range strings.Split(dir, "\n") {
		fmt.Println("Initial string: ", line)

		matches := regPhone.FindStringSubmatch(line)

		if len(matches) < 2 || matches[1] != num {
			continue
		}
		if ret != "" {
			return "Error => Too many people: " + num
		}
		fmt.Println("Current line 1: ", line)

		line = regPhone.ReplaceAllString(line, "")
		ret += "Phone => " + matches[1]

		fmt.Println("Current line 2: ", line)

		matches = regName.FindStringSubmatch(line)
		line = regName.ReplaceAllString(line, "")
		ret += ", Name => " + matches[1]

		fmt.Println("Current line 3: ", line)

		line = regSym.ReplaceAllString(line, " ")
		ret += ", Address => " + strings.TrimSpace(line)
	}
	if ret == "" {
		return "Error => Not found: " + num
	}
	return ret
}

func Phone5(dir, num string) string {

	var regPhone = regexp.MustCompile(`\+(\d{1,2}-\d{3}-\d{3}-\d{4})`)
	var regName = regexp.MustCompile(`<([^>]+)>\s?`)
	var regSym = regexp.MustCompile(`[^\-.0-9A-Za-z]+`)

	ret := ""

	for _, line := range strings.Split(dir, "\n") {
		matches := regPhone.FindStringSubmatch(line)

		if len(matches) < 2 || matches[1] != num {
			continue
		}
		fmt.Println("The state of the string: ", line)

		if ret != "" {
			return "Error => Too many people: " + num
		}

		line = regPhone.ReplaceAllString(line, "")
		ret += "Phone => " + matches[1]

		matches = regName.FindStringSubmatch(line)
		line = regName.ReplaceAllString(line, "")
		ret += ", Name => " + matches[1]

		line = regSym.ReplaceAllString(line, " ")
		ret += ", Address => " + strings.TrimSpace(line)
	}

	if ret == "" {
		return "Error => Not found: " + num
	}

	return ret
}

func phone6(dir, num string) string {

	var regPhone = regexp.MustCompile(`\+(\d{1,2}-\d{3}-\d{3}-\d{4})`)
	var regName = regexp.MustCompile(`<([^>]+)>\s?`)
	var regSym = regexp.MustCompile(`[^\-.0-9A-Za-z]+`)

	ret := ""
	for _, line := range strings.Split(dir, "\n") {
		matches := regPhone.FindStringSubmatch(line)
		if len(matches) < 2 || matches[1] != num {
			continue
		}
		if ret != "" {
			return "Error => Too many people: " + num
		}
		line = regPhone.ReplaceAllString(line, "")
		ret += "Phone => " + matches[1]

		matches = regName.FindStringSubmatch(line)
		line = regName.ReplaceAllString(line, "")
		ret += ", Name => " + matches[1]

		line := regSym.ReplaceAllString(line, " ")
		ret += ", Address => " + strings.TrimSpace(line)
	}
	if ret == "" {
		return "Error => Not Found" + num
	}
	return ret
}
