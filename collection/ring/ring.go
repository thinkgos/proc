package ring

import (
	"iter"
	"slices"
)

type Ring[T any] struct {
	head int  // 下一次写入的位置
	tail int  // 最早元素的位置
	full bool // 标记队列是否已满
	buff []T  // 缓冲区
}

// NewRing return a fixed capacity circle queue,
// if the ring exceeds the capacity, the old value will be overwritten.
func NewRing[T any](capacity int) *Ring[T] {
	if capacity < 0 {
		panic("ring: capacity must be greater than zero")
	}
	return &Ring[T]{
		head: 0,
		tail: 0,
		full: false,
		buff: make([]T, capacity),
	}
}

func (r *Ring[T]) isEmpty() bool { return !r.full && r.head == r.tail }
func (r *Ring[T]) length() int {
	if r.full {
		return len(r.buff)
	}
	if r.head >= r.tail {
		return r.head - r.tail
	}
	return len(r.buff) - r.tail + r.head
}

func (r *Ring[T]) IsFull() bool  { return r.full }
func (r *Ring[T]) IsEmpty() bool { return r.isEmpty() }
func (r *Ring[T]) Capacity() int { return len(r.buff) }
func (r *Ring[T]) Len() int      { return r.length() }

func (r *Ring[T]) Push(val T) {
	if r.full {
		r.tail = (r.tail + 1) % len(r.buff)
	}
	r.buff[r.head] = val
	r.head = (r.head + 1) % len(r.buff)
	r.full = r.head == r.tail
}

// Peek returns the oldest element from the ring.
func (r *Ring[T]) Peek() (T, bool) {
	var zero T

	if !r.full && r.head == r.tail {
		// If the ring is empty, return false
		return zero, false
	}
	return r.buff[r.tail], true
}

// PeekLatest returns the newest element from the ring.
func (r *Ring[T]) PeekLatest() (T, bool) {
	var zero T

	if r.isEmpty() {
		return zero, false
	}
	idx := (r.head - 1 + len(r.buff)) % len(r.buff)
	return r.buff[idx], true
}
func (r *Ring[T]) pop() (T, bool) {
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
func (r *Ring[T]) Pop() (T, bool) { return r.pop() }

// PopWithin removes and returns up to n elements from the ring.
// n = -1 means all elements
func (r *Ring[T]) PopWithin(n int) []T {
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
func (r *Ring[T]) Values() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range r.length() {
			idx := (r.tail + i) % len(r.buff)
			if !yield(r.buff[idx]) {
				return
			}
		}
	}
}

func (r *Ring[T]) CollectValues() []T {
	return slices.Collect(r.Values())
}
