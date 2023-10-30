package main

import (
	"fmt"
	"sync"
)

// Queue represents a queue that holds a slice
type Queue struct {
	mt          sync.RWMutex
	collections []interface{}
}

// Enqueue adds a value at the end
func (q *Queue) Enqueue(value interface{}) {
	q.mt.Lock()
	defer q.mt.Unlock()
	q.collections = append(q.collections, value)
}

// Dequeue removes a value at the front ans return the removed value
func (q *Queue) Dequeue() interface{} {
	q.mt.Lock()
	defer q.mt.Unlock()

	if q.isEmpty() {
		return nil
	}

	val := q.collections[0]
	q.collections = q.collections[1:]

	return val
}

func (q *Queue) isEmpty() bool {
	return q.Count() == 0
}

func (q *Queue) Count() int {
	return len(q.collections)
}

func (q *Queue) Collections() []interface{} {
	return q.collections
}

func useQueue() {
	queue := &Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	fmt.Println(queue.Collections())
	queue.Dequeue()
	fmt.Println(queue.Collections())
}
