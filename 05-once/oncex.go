package main

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
)

// Once 增强版Once
// 可以获取Do的error，确保能执行成功
type Once struct {
	m    sync.Mutex
	done uint32
}

func (o *Once) Do(f func() error) error {
	if atomic.LoadUint32(&o.done) == 1 {
		return nil
	}
	return o.slowDo(f)
}

func (o *Once) slowDo(f func() error) error {
	o.m.Lock()
	defer o.m.Unlock()
	var err error
	if o.done == 0 {
		err = f()
		if err == nil {
			atomic.StoreUint32(&o.done, 1)
		}
	}
	return err
}

// Done 是否初始化成功
func (o *Once) Done() bool {
	return atomic.LoadUint32(&o.done) == 1
}

func main() {
	once := Once{}
	err := once.Do(func() error {
		return errors.New("some error1 ")
	})
	fmt.Println(err)
	err = once.Do(func() error {
		return errors.New("some error2 ")
	})
	fmt.Println(err)

	fmt.Println(once.Done())
	err = once.Do(func() error {
		fmt.Println("111")
		return nil
	})
	fmt.Println(err)
	fmt.Println(once.Done())
}
