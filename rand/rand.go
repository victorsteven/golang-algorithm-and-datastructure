package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	test()
}

func test() {

	rand.Seed(time.Now().UnixNano())

	fmt.Println()
}
