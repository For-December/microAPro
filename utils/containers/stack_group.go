package containers

import (
	"sync"
	"time"
)

type StackGroup[T string | int] struct {
	stacks         map[int]*CustomStack[T]
	capacity       int
	expireDuration time.Duration
	lock           sync.Mutex
}

func NewStackGroup[T string | int](
	capacity int,
	expireDuration time.Duration,
) *StackGroup[T] {
	return &StackGroup[T]{
		stacks:         make(map[int]*CustomStack[T]),
		capacity:       capacity,
		expireDuration: expireDuration,
		lock:           sync.Mutex{},
	}
}

func (sg *StackGroup[T]) GetStack(key int) *CustomStack[T] {
	sg.lock.Lock()
	defer sg.lock.Unlock()

	if stack, ok := sg.stacks[key]; ok {
		return stack
	}

	stack := NewCustomStack[T](sg.capacity, sg.expireDuration)
	return stack
}
