package main

import "sync"

type MapConc struct {
	mp map[int]int
	mu sync.Mutex
}

func NewMap() MapConc {
	return MapConc{make(map[int]int), sync.Mutex{}}
}

func (mp *MapConc) Set(key, value int) {
	mp.mu.Lock()
	mp.mp[key] = value
	mp.mu.Unlock()
}

func (mp *MapConc) Get(key int) (int, bool) {
	mp.mu.Lock()
	defer mp.mu.Unlock()

	val, exist := mp.mp[key]
	return val, exist
}

func main() {
	mp := NewMap()

	wg := sync.WaitGroup{}

	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := range 100 {
			mp.Set(i, i)
		}
	}()
	go func() {
		defer wg.Done()
		for i := range 100 {
			mp.Set(i, i+5)
		}
	}()
	go func() {
		defer wg.Done()
		for i := range 100 {
			mp.Get(i)
		}
	}()

	wg.Wait()
}
