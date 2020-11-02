package main

import (
	"fmt"
	"sync"
)

type AtomicInt struct {
	m sync.Mutex
	n int
}

func (a *AtomicInt) Add(n int) {
	a.m.Lock()
	a.n += n
	a.m.Unlock()
}

func (a *AtomicInt) Value() int {
	a.m.Lock()
	n := a.n
	a.m.Unlock()
	return n
}

func main() {

	wait := make(chan struct{})

	var n AtomicInt

	go func() {
		n.Add(1)
		close(wait)
	}()

	n.Add(1)
	<-wait

	fmt.Println(n.Value())
}
