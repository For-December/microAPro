package containers

import (
	"sync"
	"time"
)

type StackGroup[T string | int | int64] struct {
	stacks         map[int64]*CustomStack[T]
	capacity       int
	expireDuration time.Duration
	lock           sync.Mutex
}

func NewStackGroup[T string | int | int64](
	capacity int,
	expireDuration time.Duration,
) *StackGroup[T] {
	return &StackGroup[T]{
		stacks:         make(map[int64]*CustomStack[T]),
		capacity:       capacity,
		expireDuration: expireDuration,
		lock:           sync.Mutex{},
	}
}

func (sg *StackGroup[T]) GetStack(key int64) *CustomStack[T] {
	sg.lock.Lock()
	defer sg.lock.Unlock()

	if stack, ok := sg.stacks[key]; ok {
		return stack
	}

	sg.stacks[key] = NewCustomStack[T](sg.capacity, sg.expireDuration)
	return sg.stacks[key]
}
