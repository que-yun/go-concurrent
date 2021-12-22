package main

import (
	"fmt"
)

func fanOut(ch <-chan interface{}, out []chan interface{}, async bool) {
	go func() {
		defer func() {
			for i := range out {
				close(out[i])
			}
		}()
		for v := range ch {
			v := v
			for i := 0; i < len(out); i++ {
				i := i
				if async { // 异步
					go func() {
						out[i] <- v
					}()
				} else {
					out[i] <- v
				}
			}
		}
	}()
}

func main() {
	ch := make(chan interface{}, 1)
	ch <- "aaa"
	out := make([]chan interface{}, 10)
	for i := range out {
		out[i] = make(chan interface{}, 1)
	}
	fanOut(ch, out, true)
	for i := range out {
		v := <-out[i]
		fmt.Println(v)
	}
}
