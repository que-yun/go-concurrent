package main

import (
	"fmt"
	"github.com/golang/groupcache/singleflight"
	"time"
)

func main() {
	g := singleflight.Group{}
	go func() {
		g.Do("aaa", func() (interface{}, error) {
			fmt.Println("aaaa")
			time.Sleep(1 * time.Second)
			return nil, nil
		})
	}()
	go func() {
		g.Do("bbb", func() (interface{}, error) {
			fmt.Println("bbb")
			time.Sleep(1 * time.Second)
			return nil, nil
		})
	}()
	go func() {
		g.Do("aaa", func() (interface{}, error) {
			fmt.Println("aaaa")
			time.Sleep(1 * time.Second)
			return nil, nil
		})
	}()
	go func() {
		g.Do("ccc", func() (interface{}, error) {
			fmt.Println("ccc")
			time.Sleep(1 * time.Second)
			return nil, nil
		})
	}()

	select {}
}
