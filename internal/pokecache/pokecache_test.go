package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	const interval = 5 * time.Second
	tests := []struct{
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, tc := range tests {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(tc.key, tc.val)
			val, ok := cache.Get(tc.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) != string(tc.val) {
				t.Errorf("expected to find value")
				return
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = 2 * baseTime
	cache := NewCache(baseTime)
	key := "https://example.com"
	val := []byte("testdata")
	cache.Add(key, val)

	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(waitTime)
	
	_, ok = cache.Get(key)
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}