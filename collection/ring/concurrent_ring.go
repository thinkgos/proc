package ring

import (
	"iter"
	"slices"
	"sync"
)

type ConcurrentRing[T any] struct {
	mux  sync.RWMutex
	head int  // 下一次写入的位置
	tail int  // 最早元素的位置
	full bool // 标记队列是否已满
	buff []T  // 缓冲区
}

// NewConcurrentRing return a fixed capacity concurrent circle queue,
// if the ring exceeds the capacity, the old value will be overwritten.
func NewConcurrentRing[T any](capacity int) *ConcurrentRing[T] {
	if capacity < 0 {
		panic("ring: capacity must be greater than zero")
	}
	return &ConcurrentRing[T]{
		head: 0,
		tail: 0,
		full: false,
		buff: make([]T, capacity),
	}
}

func (r *ConcurrentRing[T]) isEmpty() bool { return !r.full && r.head == r.tail }

func (r *ConcurrentRing[T]) length() int {
	if r.full {
		return len(r.buff)
	}
	if r.head >= r.tail {
		return r.head - r.tail
	}
	return len(r.buff) - r.tail + r.head
}

func (r *ConcurrentRing[T]) IsFull() bool {
	r.mux.RLock()
	defer r.mux.RUnlock()
	return r.full
}
func (r *ConcurrentRing[T]) IsEmpty() bool {
	r.mux.RLock()
	defer r.mux.RUnlock()
	return r.isEmpty()
}
func (r *ConcurrentRing[T]) Capacity() int {
	r.mux.RLock()
	defer r.mux.RUnlock()
	return len(r.buff)
}
func (r *ConcurrentRing[T]) Len() int {
	r.mux.RLock()
	defer r.mux.RUnlock()
	return r.length()
}

func (r *ConcurrentRing[T]) Push(val T) {
	r.mux.Lock()
	defer r.mux.Unlock()
	if r.full {
		r.tail = (r.tail + 1) % len(r.buff)
	}
	r.buff[r.head] = val
	r.head = (r.head + 1) % len(r.buff)
	r.full = r.head == r.tail
}

// Peek returns the oldest element from the ring.
func (r *ConcurrentRing[T]) Peek() (T, bool) {
	var zero T

	r.mux.RLock()
	defer r.mux.RUnlock()
	if !r.full && r.head == r.tail {
		// If the ring is empty, return false
		return zero, false
	}
	return r.buff[r.tail], true
}

func (r *ConcurrentRing[T]) PeekLatest() (T, bool) {
	var zero T

	r.mux.RLock()
	defer r.mux.RUnlock()
	if r.isEmpty() {
		return zero, false
	}
	idx := (r.head - 1 + len(r.buff)) % len(r.buff)
	return r.buff[idx], true
}

func (r *ConcurrentRing[T]) pop() (T, bool) {
	var zero T

	if !r.full && r.head == r.tail {
		// If the ring is empty, return false
		return zero, false
	}

	val := r.buff[r.tail]
	r.buff[r.tail] = zero // zero/nil out the obsolete elements, for GC
	r.tail = (r.tail + 1) % len(r.buff)
	r.full = false
	return val, true
}

// Pop removes and returns the oldest element from the ring.
func (r *ConcurrentRing[T]) Pop() (T, bool) {
	r.mux.Lock()
	defer r.mux.Unlock()
	return r.pop()
}

// PopWithin removes and returns up to n elements from the ring.
// n = -1 means all elements
func (r *ConcurrentRing[T]) PopWithin(n int) []T {
	r.mux.Lock()
	defer r.mux.Unlock()
	if n < 0 {
		n = r.length()
	}
	popped := make([]T, 0, n)
	for range n {
		val, ok := r.pop()
		if !ok {
			// If the queue is empty before we reach n elements, stop
			break
		}
		popped = append(popped, val)
	}
	return popped
}

// Values iterator over sequences of individual values.
func (r *ConcurrentRing[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		r.mux.RLock()
		defer r.mux.RUnlock()
		for i := range r.length() {
			idx := (r.tail + i) % len(r.buff)
			if !yield(r.buff[idx]) {
				return
			}
		}
	}
}

func (r *ConcurrentRing[T]) CollectValues() []T {
	return slices.Collect(r.Values())
}
