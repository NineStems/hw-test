package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("last push out", func(t *testing.T) {
		c := NewCache(3)
		wasInCache := c.Set("1", 100)
		require.False(t, wasInCache)
		wasInCache = c.Set("2", 200)
		require.False(t, wasInCache)
		wasInCache = c.Set("3", 300)
		require.False(t, wasInCache)
		wasInCache = c.Set("4", 400)
		require.False(t, wasInCache)
		el, inCache := c.Get("1")
		require.False(t, inCache)
		require.Nil(t, el)
	})

	t.Run("last push out with get elements", func(t *testing.T) {
		c := NewCache(3)
		wasInCache := c.Set("1", 100)
		require.False(t, wasInCache)
		wasInCache = c.Set("2", 200)
		require.False(t, wasInCache)
		wasInCache = c.Set("3", 300)
		require.False(t, wasInCache)
		_, _ = c.Get("1")
		_, _ = c.Get("2")
		_, _ = c.Get("2")
		wasInCache = c.Set("4", 400)
		require.False(t, wasInCache)
		el, inCache := c.Get("3")
		require.False(t, inCache)
		require.Nil(t, el)
	})
}

func TestCacheMultithreading(t *testing.T) {
	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}
