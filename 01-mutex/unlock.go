package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	mux := sync.Mutex{}
	go func() {
		time.Sleep(1 * time.Second)
		mux.Unlock()
		fmt.Println("B释放了A的锁")
	}()
	fmt.Println("A加锁了")
	mux.Lock()
	time.Sleep(2 * time.Second)
	mux.Unlock()
	fmt.Println("A解锁了")
}
