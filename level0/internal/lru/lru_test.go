package lru

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCache_panic(t *testing.T) {
	defer func() {
		recover()
	}()
	NewLru[int](0)

	defer func() {
		recover()
	}()
	NewLru[string](-14985)

	t.Errorf("func must panic!")
}

func TestCache_update(t *testing.T) {
	c := NewLru[int](10)

	_, ok := c.Get(13)
	require.False(t, ok)

	c.Set(13, "hello world!")
	val, ok := c.Get(13)
	require.True(t, ok)
	require.Equal(t, "hello world!", val)
}

func TestCache_eviction(t *testing.T) {
	c := NewLru[int](10)

	c.Set(100, -1)

	for i := range 20 {
		c.Set(i, i)
	}

	_, ok := c.Get(100)
	require.False(t, ok)
}

// imitation of real work our App
type app struct {
	lru Cache[int64]
}

func (a *app) get(key int64) {
	if a.lru != nil {
		_, exist := a.lru.Get(key)
		if exist {
			return
		}
		a.lru.Set(key, nil)
	}
	// imitation of query to postgres, in practice it may be more..
	time.Sleep(20 * time.Microsecond)
}

func TestComparation(t *testing.T) {
	for _, tc := range []struct {
		length   int
		capacity int
		ranged   int
	}{
		{length: 10000, capacity: 400, ranged: 1000},
		{length: 20000, capacity: 250, ranged: 1000},
		{length: 100000, capacity: 500, ranged: 1000},
	} {

		var noCache app

		tt := time.Now()
		for i := range tc.length {
			noCache.get(int64(i))
		}

		withoutCache := time.Since(tt)

		trace := make([]int64, tc.length)

		for i := range trace {
			trace[i] = rand.Int63() % int64(tc.ranged)
		}
		cached := app{lru: NewLru[int64](tc.capacity)}

		tt = time.Now()
		for _, v := range trace {
			cached.get(v)
		}
		withCache := time.Since(tt)

		t.Log("Time duraction to process", tc.length, "orders without cache", withoutCache)
		t.Log("Time duraction to process", tc.length, "orders with cache", withCache, "using", tc.capacity, "capacity of", tc.ranged, "range")

	}
}

func TestCache_concurrent(t *testing.T) {
	cache := NewLru[string](1000)
	numGoroutines := 10
	numOperationsPerGoroutine := 1000

	var wg sync.WaitGroup
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperationsPerGoroutine; j++ {
				key := "key" + string(rune(j))
				value := "value" + string(rune(id)) + string(rune(j))

				if j%2 == 0 {
					cache.Set(key, value)
				} else {
					cache.Get(key)
				}
				time.Sleep(time.Microsecond)
			}
		}(i)
	}
	wg.Wait()
}
