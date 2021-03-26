package phone_number_test

import (
	"efficient-algorithms/phone_number"
	"fmt"
	"testing"
)

func TestCreatePhoneNumber(t *testing.T) {

	nums := [10]uint{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	ans := phone_number.CreatePhoneNumber(nums)

	fmt.Println(ans)
}
