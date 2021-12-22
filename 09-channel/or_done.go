package main

import (
	"fmt"
	"reflect"
	"time"
)

/**
很多任务，只要任意一个完成就返回
*/
func or(channels ...<-chan interface{}) <-chan interface{} {
	// 特殊情况，只有零个或者1个chan
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		switch len(channels) {
		// 2个也是一种特殊情况
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			m := len(channels) / 2
			select {
			case <-or(channels[:m]...):
			case <-or(channels[m:]...):
			}
		}
	}()
	return orDone
}

// 反射实现or-done
func orByReflect(channels ...<-chan interface{}) <-chan interface{} {
	// 特殊情况，只有0个或者1个
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)
		// 利用反射构建SelectCase
		var cases []reflect.SelectCase
		for _, c := range channels {
			cases = append(cases, reflect.SelectCase{
				Dir:  reflect.SelectRecv,
				Chan: reflect.ValueOf(c),
			})
		}
		// 随机选择一个可用的case
		reflect.Select(cases)
	}()
	return orDone
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-orByReflect(
		sig(10*time.Second), sig(20*time.Second), sig(30*time.Second), sig(40*time.Second), sig(50*time.Second), sig(01*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start))
}
