package queue

import "container/list"

func NewLinkedQueue() *LinkedQueue {
	queue := &LinkedQueue{}
	queue.li = list.New()
	return queue
}

type LinkedQueue struct {
	li *list.List
}

func (t *LinkedQueue) Enqueue(v interface{}) {
	t.li.PushBack(v)
}

func (t *LinkedQueue) Dequeue() (v interface{}) {
	if t.li.Len() > 0 {
		v := t.li.Front()
		t.li.Remove(v)
		return v.Value
	}
	return nil
}

func (t *LinkedQueue) Front() (v interface{}) {
	if t.li.Len() > 0 {
		v := t.li.Front()
		return v.Value
	}
	return nil
}

func (t *LinkedQueue) Size() int {
	return t.li.Len()
}

func (t *LinkedQueue) IsEmpty() bool {
	return t.li.Len() == 0
}
