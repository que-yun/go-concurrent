package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(10)
	for i := 0; i < 100; i++ {
		wg.Done()
		test(i)
	}
	select {}
}
func test(n int) {
	wg.Add(1)
	wg.Wait()
	fmt.Println(n)
}



