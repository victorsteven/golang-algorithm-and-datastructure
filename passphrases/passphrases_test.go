package passphrases

import (
	"fmt"
	"testing"
)

func TestPlayPass(t *testing.T) {

	ans := PlayPass("MY GRANMA CAME FROM NY ON THE 23RD OF APRIL 2015", 2)

	fmt.Println("the result: ", ans)
}
