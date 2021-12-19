package main

import (
	"sync"
)

/** SliceQueue
无锁队列
*/
type SliceQueue struct {
	data []interface{}
	mux  sync.Mutex
}

func NewSliceQueue(n int) (q *SliceQueue) {
	return &SliceQueue{data: make([]interface{}, 0, n)}
}

// Enqueue 把值放在队尾
func (q *SliceQueue) Enqueue(v interface{}) {
	q.mux.Lock()
	q.data = append(q.data, v)
	q.mux.Unlock()
}

// Dequeue 移去队头并返回
func (q *SliceQueue) Dequeue() interface{} {
	q.mux.Lock()
	if len(q.data) == 0 {
		q.mux.Unlock()
		return nil
	}
	v := q.data[0]
	q.data = q.data[1:]
	q.mux.Unlock()
	return v
}
