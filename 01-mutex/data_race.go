package main

import (
	"fmt"
	"sync"
)

/**
数据竞争问题:
	多个goroutine 同时操作一个共享变量
	会导致data race   使用 go run -race data_race.go 可以监测
使用sync.Mutex 的lock和unlock解决 data race问题

	go tool compile -S data_race.go   获取汇编代码
*/

func main() {
	// 数据竞争
	dataRace()
	// 加锁
	useLock()
	// 内嵌锁
	useLockByStruct()

	channel()
}

/**
1. data race
*/
func dataRace() {
	var count = 0
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}

/**
2. 使用锁
*/
func useLock() {
	// 互斥锁保护计数器
	var mu sync.Mutex
	var count = 0
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}

/**
3. 结构体内嵌锁
*/

type Counter struct {
	sync.Mutex
	Count uint64
}

func useLockByStruct() {
	var counter Counter
	// 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				counter.Lock()
				counter.Count++
				counter.Unlock()
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(counter.Count)
}

/**
4. channel 实现计数
 */
func channel() {
	ch := make(chan struct{})
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				ch <- struct{}{}
			}
		}()
	}

	go func() {
		wg.Wait() // 等待上面所有的 goroutine 运行完成
		close(ch) // 关闭ch通道
	}()

	count := 0
	for range ch { // 如果ch通道读取完了(ch是关闭状态), 则for循环结束
		count++
	}
	fmt.Println("count的值是:", count)
}
