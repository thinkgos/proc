package queue

// element is an element of the Queue implement with list.
type element[T comparable] struct {
	next  *element[T]
	value T
}

// Queue represents a singly linked list.
type Queue[T comparable] struct {
	head   *element[T]
	tail   *element[T]
	length int
}

// New creates a Queue. which implement queue.Interface.
func New[T comparable]() *Queue[T] {
	return new(Queue[T])
}

// Len returns the length of this queue.
func (q *Queue[T]) Len() int { return q.length }

// IsEmpty returns true if this Queue contains no elements.
func (q *Queue[T]) IsEmpty() bool { return q.Len() == 0 }

// Clear initializes or clears queue.
func (q *Queue[T]) Clear() { q.head, q.tail, q.length = nil, nil, 0 }

// Add items to the queue.
func (q *Queue[T]) Add(v T) {
	e := &element[T]{value: v}
	if q.tail == nil {
		q.head, q.tail = e, e
	} else {
		q.tail.next = e
		q.tail = e
	}
	q.length++
}

// Peek retrieves, but does not remove, the head of this Queue, or return nil if this Queue is empty.
func (q *Queue[T]) Peek() (v T, ok bool) {
	if q.head != nil {
		return q.head.value, true
	}
	return v, false
}

// Poll retrieves and removes the head of the this Queue, or return nil if this Queue is empty.
func (q *Queue[T]) Poll() (v T, ok bool) {
	if q.head != nil {
		v = q.head.value
		q.head = q.head.next
		if q.head == nil {
			q.tail = nil
		}
		q.length--
		ok = true
	}
	return v, ok
}

// Contains returns true if this queue contains the specified element.
func (q *Queue[T]) Contains(val T) bool {
	for e := q.head; e != nil; e = e.next {
		if val == e.value {
			return true
		}
	}
	return false
}

// Remove a single instance of the specified element from this queue, if it is present.
func (q *Queue[T]) Remove(val T) {
	for pre, e := q.head, q.head; e != nil; {
		if val == e.value {
			switch {
			case q.head == e && q.tail == e:
				q.head, q.tail = nil, nil
			case q.head == e:
				q.head = e.next
			case q.tail == e:
				q.tail = pre
				q.tail.next = nil
			default:
				pre.next = e.next
			}
			e.next = nil
			q.length--
			return
		}
		pre = e
		e = e.next
	}
}
