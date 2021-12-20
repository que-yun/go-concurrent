package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	pool := sync.Pool{
		New: func() interface{} {
			return new(bytes.Buffer)
		},
	}
	service := pool.Get()
	if service == nil {
		service = new(bytes.Buffer)
	}
	buffer := service.(*bytes.Buffer)
	buffer.WriteString("xxx")
	fmt.Println(buffer)
	pool.Put(service)
}
