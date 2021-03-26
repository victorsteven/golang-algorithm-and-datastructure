package IP_Validation

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidIP(t *testing.T) {

	fixtures := []struct {
		IP      string
		IsValid bool
	}{
		{
			"1.2.3.4",
			true,
		},
		{
			"123.45.67.89",
			true,
		},
		{
			"1.2.3",
			false,
		},
		{
			"1.2.3.4.5",
			false,
		},
		{
			"123.456.78.90",
			false,
		},
		{
			"123.045.067.089",
			false,
		},
	}

	for _, v := range fixtures {

		t.Run(fmt.Sprintf("%s test", v.IP), func(t *testing.T) {
			isValid := Is_valid_ip2(v.IP)
			assert.Equal(t, v.IsValid, isValid)
		})

	}
}
