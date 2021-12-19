package main

import (
	"fmt"
	"sync"
	"time"
)

// 线程安全的计数器
type Counter struct {
	mux   sync.Mutex
	count uint64
}

// 对计数值加一
func (c *Counter) Incr() {
	c.mux.Lock()
	c.count++
	c.mux.Unlock()
}

// 获取当前的计数值
func (c *Counter) Count() uint64 {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.count
}

// sleep 1秒， 然后计数值加1
func worker(c *Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second)
	c.Incr()
}

func main() {
	var counter Counter
	var wg sync.WaitGroup
	// WaitGroup的值设置为10
	wg.Add(10)
	// 启动10个goroutine执行++
	for i := 0; i < 10; i++ {
		go worker(&counter, &wg)
	}
	// 等待goroutine都完成
	wg.Wait()
	// 输出当前计数器的值
	fmt.Println(counter.Count())
}
