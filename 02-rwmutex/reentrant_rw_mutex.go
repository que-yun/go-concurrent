package main

import (
	"fmt"
	"github.com/petermattis/goid"
	"sync"
	"sync/atomic"
	"time"
)

/**
TODO： 未完成
*/

// ReentrantRWMutex 可重入读写锁
type ReentrantRWMutex struct {
	sync.RWMutex
	owner     int64 // 当前持有锁的goroutine id
	recursion int32 // 这个goroutine 重入的次数
}

func (m *ReentrantRWMutex) Lock() {
	gid := goid.Get()
	// 如果当前持有锁的goroutine就是这次调用的goroutine，说明是重入
	if atomic.LoadInt64(&m.owner) == gid {
		m.recursion++
		return
	}

	m.RWMutex.Lock()
	// 获得锁的goroutine第一次调用，记录下它的goroutine id， 调用次数加1
	atomic.StoreInt64(&m.owner, gid)
	m.recursion = 1
}

func (m *ReentrantRWMutex) Unlock() {
	gid := goid.Get()
	// 非持有锁的goroutine尝试释放锁，错误的使用
	if atomic.LoadInt64(&m.owner) != gid {
		panic(fmt.Sprintf("wrong the owner(%d): %d!", m.owner, gid))
	}
	// 调用次数减1
	m.recursion--
	// 如果这个goroutine还没有完全释放，则直接返回
	if m.recursion != 0 {
		return
	}
	// 此goroutine最后一次调用，需要释放锁
	atomic.StoreInt64(&m.owner, -1)
	m.RWMutex.Unlock()
}

func main() {
	rwMutex := ReentrantRWMutex{}
	rwMutex.Lock()
	defer rwMutex.Unlock()
	time.Sleep(1 * time.Second)
}
