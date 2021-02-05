package IP_Validation

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// Write an algorithm that will identify valid IPv4 addresses in dot-decimal format. IPs should be considered valid if they consist of four octets, with values between 0 and 255, inclusive.
//
//Input to the function is guaranteed to be a single string.
//
//Examples
//Valid inputs:
//
//1.2.3.4
//123.45.67.89
//Invalid inputs:
//
//1.2.3
//1.2.3.4.5
//123.456.78.90
//123.045.067.089
//Note that leading zeros (e.g. 01.02.03.04) are considered invalid.

func Is_valid_ip1(ip string) bool {

	res := net.ParseIP(ip)
	if res == nil {
		return false
	}
	return true
}

func split(value string, separator string) []string {

	var parts []string
	last := 0

	for i, v := range value {
		if string(v) == separator {
			parts = append(parts, value[last:i])
			last = i + 1
		}
	}
	parts = append(parts, value[last:])

	fmt.Println("the parts returned: ", parts)

	return parts
}

func Is_valid_ip(ip string) bool {

	//parts := split(ip, ".")
	//if len(parts) != 4 {
	//	return false
	//}

	//could simply use split string:
	parts := strings.Split(ip, ".")
	if len(parts) != 4 {
		return false
	}

	for _, value := range parts {
		if value[0] == '0' && len(value) > 1 {
			return false
		}

		num, err := strconv.Atoi(value)
		if err != nil {
			return false
		}

		if num > 255 || num < 0 {
			return false
		}
	}

	return true
}



