package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(GoID())
	go func() {
		fmt.Println(GoID())
	}()
	time.Sleep(500 * time.Millisecond)
}

// GoID 获取goroutineId
func GoID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	fmt.Println(string(buf[:n]))
	// 得到Id字符串
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
