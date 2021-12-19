package main

import (
	"fmt"
	"sync"
)

func main() {
	var once sync.Once
	// 第一个
	once.Do(func() {
		fmt.Println("1")
	})
	// 第二个
	once.Do(func() {
		fmt.Println("2")
	})
}
