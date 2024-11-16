package containers

import (
	"container/list"
	"microAPro/utils/logger"
	"sync"
	"time"
)

type Message[T string | int | int64] struct {
	Value    T
	ExpireAt time.Time
}

type CustomStack[T string | int | int64] struct {
	capacity       int
	expireDuration time.Duration
	stack          *list.List
	lock           sync.Mutex
}

func NewCustomStack[T string | int | int64](capacity int,
	expireDuration time.Duration) *CustomStack[T] {
	return &CustomStack[T]{
		capacity:       capacity,
		expireDuration: expireDuration,
		stack:          list.New(),
		lock:           sync.Mutex{},
	}
}

func (cs *CustomStack[T]) Push(value T) {
	cs.lock.Lock()
	defer cs.lock.Unlock()

	// 如果已满，移除最底部的消息（最早进入的）
	for cs.stack.Len() >= cs.capacity {
		bottom := cs.stack.Front()

		// 肯定不为空
		if bottom != nil {
			cs.stack.Remove(bottom)
		} else {
			logger.Error("断言出错：", cs)
			return
		}
	}

	cs.stack.PushBack(Message[T]{
		Value:    value,
		ExpireAt: time.Now().Add(cs.expireDuration),
	})
}

func (cs *CustomStack[T]) Pop() (res T, ok bool) {
	cs.lock.Lock()
	defer cs.lock.Unlock()

	if cs.stack.Len() == 0 {
		return
	}

	back := cs.stack.Back()
	if back == nil {
		logger.Error("断言出错，不会为空：", cs)
		return
	}

	message := back.Value.(Message[T])

	// 过期的消息不再返回，并清除所有过期消息
	if time.Now().After(message.ExpireAt) {
		cs.stack.Remove(back)

		// 从后往前清除所有过期消息
		for el := cs.stack.Back(); el != nil; el = cs.stack.Back() {
			if el.Value.(Message[T]).ExpireAt.After(time.Now()) {
				logger.Error("未过期却被删除？", cs)
			}
			cs.stack.Remove(el)
		}
		return
	}

	cs.stack.Remove(back)
	return message.Value, true
}
