package main

import (
	"flag"
	"fmt"
)

var ip string
var port int

func init() {
	flag.StringVar(&ip, "ip", "0.0.0.0", "ip address")
	flag.IntVar(&port, "port", 8000, "port number")
}

func main() {
	flag.Parse()
	fmt.Printf("%s:%d", ip, port)
}
