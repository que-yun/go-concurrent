package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// unlockByOther()
	// copyLock()
	l := &sync.Mutex{}
	foo1(l)
}

/**
3. 不可重入
*/
func foo1(l sync.Locker) {
	fmt.Println("in foo")
	l.Lock()
	bar(l)
	l.Unlock()
}

func bar(l sync.Locker) {
	l.Lock()
	fmt.Println("in bar")
	l.Unlock()
}

/**
2. 静态检查
*/
func copyLock() {
	var c Counter
	c.Lock()
	defer c.Unlock()
	c.Count++
	foo(c)
}

func foo(c Counter) {
	c.Lock()
	defer c.Unlock()
	fmt.Println("in foo")
}

/**
1. 被其他人解锁了
*/
func unlockByOther() {
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
