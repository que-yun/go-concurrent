package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var counter Counter
	for i := 0; i < 10; i++ {
		go func() {
			for {
				fmt.Println(counter.Count())
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for {
		counter.Incr()
		time.Sleep(time.Second)
	}
}

// Counter 一个线程安全的计数器
type Counter struct {
	mux   sync.RWMutex
	count uint64
}

// 使用写锁保护
func (c *Counter) Incr() {
	c.mux.Lock()
	c.count++
	c.mux.Unlock()
}

// 使用读锁保护
func (c *Counter) Count() uint64 {
	c.mux.RLock()
	defer c.mux.RUnlock()
	return c.count
}
