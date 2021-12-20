package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	var i int32 = 2
	atomic.AddInt32(&i, 20)
	fmt.Println(i)
	var u uint32 = 22
	// uint32 实现减法  -20
	atomic.AddUint32(&u, ^uint32(20-1))
	fmt.Println(u)
}
