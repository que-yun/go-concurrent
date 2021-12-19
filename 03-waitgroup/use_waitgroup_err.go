package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// addError()
	firstAdd()
}

func addError() {
	var wg sync.WaitGroup
	wg.Add(10)
	wg.Add(-10)
	wg.Done()
}

/**
2. wg.Add 需要提前在当前goroutine设置,否则执行结果不可预料
*/
func firstAdd() {
	var wg sync.WaitGroup
	go dosomething(100, &wg) // 启动第一个goroutine
	go dosomething(110, &wg) // 启动第二个goroutine
	go dosomething(120, &wg) // 启动第三个goroutine
	go dosomething(130, &wg) // 启动第四个goroutine

	wg.Wait() // 主goroutine等待完成
	fmt.Println("Done")
}

func dosomething(millisecs time.Duration, wg *sync.WaitGroup) {
	//duration := millisecs * time.Millisecond
	//time.Sleep(duration) // 故意sleep一段时间

	wg.Add(1)
	fmt.Println("后台执行, duration:", "duration")
	wg.Done()
}
