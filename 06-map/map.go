package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	m.Store("hello", "world")
	if load, ok := m.Load("hello"); ok {
		fmt.Println(load)
	}
	m.Delete("hello")
}
