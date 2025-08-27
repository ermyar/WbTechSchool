package lru

import (
	"container/list"
)

type Cache[T comparable] interface {
	Get(T) (interface{}, bool)

	Set(T, interface{}) error

	Capacity() int

	Clear()
}

type Lru[T comparable] struct {
	mp       map[T]*list.Element
	list     *list.List
	capacity int
}

type kvPair[T comparable] struct {
	key   T
	value interface{}
}

func (lru *Lru[T]) pushTop(key T, value interface{}) {
	if ptr, exist := lru.mp[key]; exist {
		lru.list.MoveToFront(ptr)
	} else if len(lru.mp) < lru.capacity {
		ptr := lru.list.PushFront(kvPair[T]{key, value})
		lru.mp[key] = ptr
	} else {
		ptr := lru.list.PushFront(kvPair[T]{key, value})
		lru.mp[key] = ptr
		pair := (lru.list.Back().Value).(kvPair[T])
		lru.list.Remove(lru.list.Back())
		delete(lru.mp, pair.key)
	}
}

func (lru *Lru[T]) Get(key T) (interface{}, bool) {
	if ptr, exist := lru.mp[key]; exist {
		value := ptr.Value
		lru.pushTop(key, value)
		return value.(kvPair[T]).value, true
	}
	return nil, false
}

func (lru *Lru[T]) Set(key T, value interface{}) error {
	lru.pushTop(key, value)
	return nil
}

func (lru *Lru[T]) Clear() {
	lru.list.Init()
	for k := range lru.mp {
		delete(lru.mp, k)
	}
}

func (lru *Lru[T]) Capacity() int {
	return lru.capacity
}

func NewLru[T comparable](capacity int) Cache[T] {
	if capacity <= 0 {
		panic("non-positive capacity")
	}
	return &Lru[T]{
		make(map[T]*list.Element),
		list.New(),
		capacity,
	}
}
