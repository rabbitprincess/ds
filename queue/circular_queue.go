package queue

import "fmt"

func NewCircularQueue(_cap int64) *CircularQueue {
	if _cap == 0 {
		return nil
	}
	return &CircularQueue{
		Data: make([]interface{}, _cap),
		cap:  _cap,
		head: 0,
		tail: 0,
	}
}

//---------------------------------------------------------------------------------------------------------------------//
// Circular Queue

type CircularQueue struct {
	Data []interface{}
	cap  int64
	head int64
	tail int64
}

func (t *CircularQueue) Cap() int64 {
	return t.cap
}

func (t *CircularQueue) Head() interface{} {
	return t.Data[t.head]
}

func (t *CircularQueue) IsEmpty() bool {
	if t.head == t.tail {
		return true
	}
	return false
}

func (t *CircularQueue) IsFull() bool {
	if t.head == (t.tail+1)%t.cap {
		return true
	}
	return false
}

func (t *CircularQueue) Enqueue(_i_data interface{}) error {
	if t.IsFull() == true {
		return fmt.Errorf("Queue is full | cap - %d", t.cap)
	}
	t.Data[t.tail] = _i_data
	t.tail = (t.tail + 1) % t.cap
	return nil
}

func (t *CircularQueue) Dequeue() (i_data interface{}) {
	if t.IsEmpty() {
		return nil
	}
	i_data = t.Data[t.head]
	t.head = (t.head + 1) % t.cap
	return i_data
}
