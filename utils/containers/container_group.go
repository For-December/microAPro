package containers

import (
	"sync"
	"time"
)

type ContainerGroup[T interface{}] struct {
	stacks         map[int64]*T
	capacity       int
	initFunc       func() *T // 初始化容器的函数
	expireDuration time.Duration

	lock sync.Mutex
}

func NewContainerGroup[T interface{}](
	capacity int,
	initFunc func() *T, // 初始化容器的函数
) *ContainerGroup[T] {
	return &ContainerGroup[T]{
		stacks:   make(map[int64]*T),
		initFunc: initFunc,
		lock:     sync.Mutex{},
	}
}

func (sg *ContainerGroup[T]) GetContainer(key int64) *T {
	sg.lock.Lock()
	defer sg.lock.Unlock()

	if content, ok := sg.stacks[key]; ok {
		return content
	}

	sg.stacks[key] = sg.initFunc()
	return sg.stacks[key]
}
