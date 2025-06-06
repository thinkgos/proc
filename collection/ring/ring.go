package ring

import "iter"

type Ring[T any] struct {
	head int  // 下一次写入的位置
	tail int  // 最早元素的位置
	full bool // 标记队列是否已满
	buff []T  // 缓冲区
}

func New[T any](capacity int) *Ring[T] {
	if capacity < 0 {
		panic("ring: size must be greater than zero")
	}
	return &Ring[T]{
		head: 0,
		tail: 0,
		full: false,
		buff: make([]T, capacity),
	}
}

func (r *Ring[T]) IsFull() bool  { return r.full }
func (r *Ring[T]) Capacity() int { return len(r.buff) }

func (r *Ring[T]) Len() int {
	if r.full {
		return len(r.buff)
	}
	return r.head
}

func (r *Ring[T]) Push(val T) {
	if r.full {
		r.tail = (r.tail + 1) % len(r.buff)
	}

	r.buff[r.head] = val
	r.head = (r.head + 1) % len(r.buff)
	r.full = r.head == r.tail
}

// All iterator over sequences of individual values.
func (r *Ring[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := range r.Len() {
			idx := (r.tail + i) % len(r.buff)
			if !yield(r.buff[idx]) {
				return
			}
		}
	}
}
